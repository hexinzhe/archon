import datetime
import base64
import hashlib
import math

def _write_cpp_header(f, include_guard):
    year = datetime.date.today().year
    f.write((
        "// Copyright %(year)d The Chromium Authors. All rights reserved.\n"
        "// Use of this source code is governed by a BSD-style license"
        " that can be\n"
        "// found in the LICENSE file.\n\n"
        "// This file is generated by print_log_list.py\n"
        "#ifndef %(include_guard)s\n"
        "#define %(include_guard)s\n\n"
        "#include <stddef.h>\n\n" %
        {"year": year, "include_guard": include_guard}))


def _write_log_info_struct_definition(f):
    f.write(
        "struct CTLogInfo {\n"
        "  const char* const log_key;\n"
        "  const size_t log_key_length;\n"
        "  const char* const log_name;\n"
        "  const char* const log_url;\n"
        "};\n\n")


def split_and_hexify_binary_data(bin_data):
    hex_data = "".join(["\\x%.2x" % ord(c) for c in bin_data])
    # line_width % 4 must be 0 to avoid splitting the hex-encoded data
    # across '\' which will escape the quotation marks.
    line_width = 68
    assert line_width % 4 == 0
    num_splits = int(math.ceil(len(hex_data) / float(line_width)))
    return ['"%s"' % hex_data[i * line_width:(i + 1) * line_width]
            for i in range(num_splits)]


def write_log_ids_array(f, log_ids, array_name):
    num_logs = len(log_ids)
    log_ids.sort()
    log_id_length = len(log_ids[0]) + 1
    f.write("// The list is sorted.\n")
    f.write("const char %s[][%d] = {\n" % (
            array_name, log_id_length))
    for i in range(num_logs):
        f.write("    ")
        split_hex_id = split_and_hexify_binary_data(log_ids[i])
        f.write("\n    ".join(split_hex_id))
        if (i < num_logs - 1):
            f.write(',\n')
    f.write('};\n\n')


def _write_cpp_footer(f, include_guard):
    f.write("\n#endif  // %s\n" % include_guard)


def _find_google_operator_id(json_log_list):
    goog_operator = filter(
        lambda op: op["name"] == "Google", json_log_list["operators"])
    assert(len(goog_operator) == 1)
    return goog_operator[0]["id"]


def generate_cpp_header(json_log_list, output_file):
    """Generate a header file of known logs to be included by Chromium."""
    include_guard = (output_file.upper().replace('.', '_').replace('/', '_')
                     + '_')
    f = open(output_file, "w")
    _write_cpp_header(f, include_guard)
    logs = json_log_list["logs"]
    list_code = []
    google_operator_id = _find_google_operator_id(json_log_list)
    google_log_ids = []
    for log in logs:
        log_key = base64.decodestring(log["key"])
        split_hex_key = split_and_hexify_binary_data(log_key)
        s = "    {"
        s += "\n     ".join(split_hex_key)
        s += ',\n     %d' % (len(log_key))
        s += ',\n     "%s"' % (log["description"])
        s += ',\n     "https://%s/"}' % (log["url"])
        list_code.append(s)

        # operated_by is a list, in practice we have not witnessed
        # a log co-operated by more than one operator. Ensure we take this
        # case into consideration if it ever happens.
        assert(len(log["operated_by"]) == 1)
        operator_id = log["operated_by"][0]
        if operator_id == google_operator_id:
            google_log_ids.append(hashlib.sha256(log_key).digest())

    _write_log_info_struct_definition(f)
    f.write("const CTLogInfo kCTLogList[] = {\n")
    f.write(",\n" . join(list_code))
    f.write("};\n")
    f.write("\nconst size_t kNumKnownCTLogs = %d;\n\n" % len(logs))

    write_log_ids_array(f, google_log_ids, 'kGoogleLogIDs')
    f.write("const size_t kNumGoogleLogs = %d;\n" % len(google_log_ids))

    _write_cpp_footer(f, include_guard)

