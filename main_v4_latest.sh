#!/bin/sh

version=$1
language=$2
command=$3
value=$4
source_prefix=$5
source=$6

sh_version="4.4.4"
old_ifs=$IFS
base_structure_dir="_base_structure"
runtime_dir="$(dirname "$(readlink -f "$0")")"
runtime_base_structure_dir="$runtime_dir/$base_structure_dir"
runtime_file=$(basename "$0")
latest_raw_url="https://raw.githubusercontent.com/dalikewara/ayapingping-sh/master/main_v4.sh"
latest_file="main_v4_latest.sh"
latest_output_filepath="$runtime_dir/$latest_file"
current_dir=$(pwd)
name="AyaPingPing"
language_golang="Golang"
language_golang_command="ayapingping-go"
language_golang_feature_example="feature1,feature2"
language_golang_domain_example="domain1.go,domain2.go"
language_golang_common_example="commonFunction1.go,commonFunction2.go"
language_python="Python"
language_python_command="ayapingping-py"
language_python_feature_example="feature1,feature2"
language_python_domain_example="domain_1.py,domain_2.py"
language_python_common_example="common_function_1.py,common_function_2.py"
language_typescript="TypeScript"
language_typescript_command="ayapingping-ts"
language_typescript_feature_example="feature1,feature2"
language_typescript_domain_example="domain1.ts,domain2.ts"
language_typescript_common_example="commonFunction1.ts,commonFunction2.ts"
import_feature_command="importFeature"
export_feature_command="exportFeature"
import_domain_command="importDomain"
export_domain_command="exportDomain"
import_common_command="importCommon"
export_common_command="exportCommon"
version_command="version"
import_command_source_prefix="from"
export_command_source_prefix="to"
import_or_export_source_example="/path/to/local/project | https://example.com/user/project.git | git@example.com:user/project.git"
feature_dir="features"
domain_dir="domain"
common_dir="common"
dependency_feature_key="features"
dependency_domain_key="domains"
dependency_common_key="commons"
dependency_external_key="externals"
value_delimiter=","

read_input_value=""

trap_is_ok=false
trap_project_name=""
trap_path=""

# last function variable prefix: _9

main() {
  if is_latest; then
    trap cleanup EXIT

    generator

    exit 0
  fi

  wget -q --no-cache -O "$latest_output_filepath" "$latest_raw_url" || true

  if [ "$(head -c 9 "$latest_output_filepath")" != "#!/bin/sh" ]; then
    curl -s -o "$latest_output_filepath" "$latest_raw_url" || true
  fi

  if is_file "$latest_output_filepath" && [ "$(head -c 9 "$latest_output_filepath")" = "#!/bin/sh" ]; then
    chmod +x $latest_output_filepath

    $latest_output_filepath "$version" "$language" "$command" "$value" "$source_prefix" "$source"

    exit 0
  else
    trap cleanup EXIT

    generator

    exit 0
  fi
}

generator() {
  print_version

  if is_version; then
    return 0
  fi

  echo ""

  if is_command_all_empty; then
    create_new_project
  elif is_import_or_export; then
    import_or_export_something
  else
    print_invalid_command_warning
  fi
}

create_new_project() {
  read_project_name
  _0_project_name="$read_input_value"
  read_go_module
  _0_go_module="$read_input_value"
  read_dont_use_venv
  _0_dont_use_venv="$read_input_value"

  echo ""
  echo "Project name: $_0_project_name"

  if is_golang; then
    echo "Go Module: $_0_go_module"
  elif is_python; then
    if [ "$_0_dont_use_venv" = true ]; then
      echo "Using virtual environment: no"
    else
      echo "Using virtual environment: yes"
    fi
  fi

  echo ""

  create_dir "$_0_project_name"

  trap_project_name="$_0_project_name"

  copy_contents "$runtime_base_structure_dir" "$_0_project_name"

  replace_go_module_in_directory_files "$_0_project_name" "$(get_go_module_from_path "$runtime_dir")/$base_structure_dir" "$_0_go_module"

  remove_unwanted_contents_from_directory "$_0_project_name"

  create_env "$_0_project_name"

  post_config "$_0_project_name" "$_0_go_module" "$_0_dont_use_venv"

  trap_is_ok=true
}

create_env() {
  if [ "$#" -ne 1 ]; then
    echo "[create_env] syntax error... usage: $0 <base_dir>"
    exit 1
  fi

  if is_file "$1/.env.example" && ! is_dir "$1/.env.example"; then
    copy_file "$1/.env.example" "$1/.env"
  elif ! is_file "$1/.env" && ! is_dir "$1/.env"; then
    touch "$1/.env"
  fi

  echo "[create_env] '.env'... created to '$1' (existing file will not be overwritten)"
}

import_or_export_something() {
  validate_import_or_export_command

  _7_source_dir=$(get_actual_source_dir)
  _7_dest_dir=$(get_actual_dest_dir)
  _7_dest_go_module=$(get_go_module_from_path "$_7_dest_dir")

  if is_import_or_export_feature; then
    copy_dependency_feature "$value" "$_7_source_dir" "$_7_dest_dir"
  elif is_import_or_export_domain; then
    copy_dependency_domain "$value" "$_7_source_dir" "$_7_dest_dir"
  elif is_import_or_export_common; then
    copy_dependency_common "$value" "$_7_source_dir" "$_7_dest_dir"
  fi

  post_config "$_7_dest_dir" "$_7_dest_go_module" "true"

  trap_is_ok=true
}

cleanup() {
  IFS=$old_ifs

  if is_version; then
    return 0
  fi

  if [ "$trap_is_ok" = true ]; then
    echo "[process] status... everything is ok"
    echo "[cleanup] cleaning up... running"

    remove_path_values "$trap_path"

    echo "[cleanup] cleaning up... done"
    echo "done"
  else
    echo "[process] status... aborted"
    echo "[cleanup] cleaning up... running"

    remove_path_values "$trap_project_name"
    remove_path_values "$trap_path"

    echo "[cleanup] cleaning up... done"
    echo "aborted"
  fi
}

is_latest() {
  [ "$runtime_file" = "$latest_file" ]
}

is_golang() {
  [ "$language" = "$language_golang" ]
}

is_python() {
  [ "$language" = "$language_python" ]
}

is_typescript() {
  [ "$language" = "$language_typescript" ]
}

is_import_feature() {
  [ "$command" = "$import_feature_command" ] && [ "$source_prefix" = "$import_command_source_prefix" ]
}

is_export_feature() {
  [ "$command" = "$export_feature_command" ] && [ "$source_prefix" = "$export_command_source_prefix" ]
}

is_import_domain() {
  [ "$command" = "$import_domain_command" ] && [ "$source_prefix" = "$import_command_source_prefix" ]
}

is_export_domain() {
  [ "$command" = "$export_domain_command" ] && [ "$source_prefix" = "$export_command_source_prefix" ]
}

is_import_common() {
  [ "$command" = "$import_common_command" ] && [ "$source_prefix" = "$import_command_source_prefix" ]
}

is_export_common() {
  [ "$command" = "$export_common_command" ] && [ "$source_prefix" = "$export_command_source_prefix" ]
}

is_version() {
  [ "$command" = "$version_command" ] && [ "$value" = "" ] && [ "$source_prefix" = "" ] && [ "$source" = "" ]
}

is_import () {
  is_import_feature || is_import_domain || is_import_common
}

is_export() {
  is_export_feature || is_export_domain || is_export_common
}

is_import_or_export() {
  is_import || is_export
}

is_import_or_export_feature() {
  is_import_feature || is_export_feature
}

is_import_or_export_domain() {
  is_import_domain || is_export_domain
}

is_import_or_export_common() {
  is_import_common || is_export_common
}

is_command_all_empty() {
  [ "$command" = "" ] && [ "$value" = "" ] && [ "$source_prefix" = "" ] && [ "$source" = "" ]
}

is_command_valid() {
  [ "$command" != "" ] && [ "$value" != "" ] && [ "$source_prefix" != "" ] && [ "$source" != "" ]
}

is_import_or_export_command_valid() {
  is_import_or_export && is_command_valid
}

is_file() {
  if [ "$#" -ne 1 ]; then
    echo "[is_file] syntax error... usage: $0 <filepath>"
    exit 1
  fi

  [ -f "$1" ]
}

is_dir() {
  if [ "$#" -ne 1 ]; then
    echo "[is_dir] syntax error... usage: $0 <dir>"
    exit 1
  fi

  [ -d "$1" ]
}

is_git_url() {
  if [ "$#" -ne 1 ]; then
    echo "[is_git_url] syntax error... usage: $0 <url>"
    exit 1
  fi

  if [ -z "$1" ]; then
    return 1
  fi

  if [ "$(echo "$1" | grep -E '\.git$' | grep -E 'https:\/\/|http:\/\/|git@')" ]; then
    return 0
  else
    return 1
  fi
}

get_feature_example() {
  if is_golang; then
    echo "$language_golang_feature_example"
  elif is_python; then
    echo "$language_python_feature_example"
  elif is_typescript; then
    echo "$language_typescript_feature_example"
  else
    echo "feature1,feature2"
  fi
}

get_domain_example() {
  if is_golang; then
    echo "$language_golang_domain_example"
  elif is_python; then
    echo "$language_python_domain_example"
  elif is_typescript; then
    echo "$language_typescript_domain_example"
  else
    echo "domain1.extension,domain2.extension"
  fi
}

get_common_example() {
  if is_golang; then
    echo "$language_golang_common_example"
  elif is_python; then
    echo "$language_python_common_example"
  elif is_typescript; then
    echo "$language_typescript_common_example"
  else
    echo "commonFunction1.extension,commonFunction2.extension"
  fi
}

get_language_command_example() {
  if is_golang; then
    echo "$language_golang_command"
  elif is_python; then
    echo "$language_python_command"
  elif is_typescript; then
    echo "$language_typescript_command"
  else
    echo "ayapingping-[language]"
  fi
}

get_command_example() {
  echo "$import_feature_command | $import_domain_command | $import_common_command | $export_feature_command | $export_domain_command | $export_common_command"
}

get_value_example() {
  if is_import_feature || is_export_feature; then
    get_feature_example
  elif is_import_domain || is_export_domain; then
    get_domain_example
  elif is_import_common || is_export_common; then
    get_common_example
  else
    echo "value1,value2"
  fi
}

get_source_prefix_example() {
  echo "$import_command_source_prefix | $export_command_source_prefix"
}

get_source_example() {
  echo "$import_or_export_source_example"
}

get_source_dir() {
  if is_git_url "$source"; then
    get_source_dir_from_git
    return 0
  fi

  echo "$source"
}

get_source_dir_from_git() {
  create_dir "tmp"

  _1_git_source_dir="tmp/tmp-git-project-$(date +%s%N)"

  trap_path="$trap_path,$_1_git_source_dir"

  git clone "$source" "$_1_git_source_dir" || exit 1

  echo "$_1_git_source_dir"
}

get_actual_source_dir() {
  if is_export; then
    echo "$current_dir"
    return 0
  fi

  get_source_dir
}

get_actual_dest_dir() {
  if is_export; then
    get_source_dir
    return 0
  fi

  echo "$current_dir"
}

get_go_module_from_path() {
  if [ "$#" -ne 1 ]; then
    echo "[get_go_module_from_path] syntax error... usage: $0 <path>"
    exit 1
  fi

  if ! is_golang; then
    echo ""
    return 1
  fi

  cd "$1" || return 1

  go list -m -modfile go.mod

  cd "$current_dir" || return 1
}

get_raw_json_from_file() {
  if [ "$#" -ne 1 ]; then
    echo "[get_raw_json_from_file] syntax error... usage: $0 <filepath>"
    exit 1
  fi

  cat "$1" | tr -d '[:space:]'
}

get_raw_json_from_file_for_external() {
  if [ "$#" -ne 1 ]; then
    echo "[get_raw_json_from_file_for_external] syntax error... usage: $0 <filepath>"
    exit 1
  fi

  cat "$1" | tr -d '\n' | tr -s '[:blank:]' ' ' | sed 's/* //g'
}

get_json_value_by_key() {
  if [ "$#" -ne 2 ]; then
    echo "[get_json_value_by_key] syntax error... usage: $0 <raw_json> <key>"
    exit 1
  fi

  get_clean_json_value "$(echo "$1" | tr -d '[:space:]' | sed -n -e "s/.*\"$2\":\[\([^]]*\)\].*/\1/p")"
}

get_json_value_by_key_for_external() {
  if [ "$#" -ne 2 ]; then
    echo "[get_json_value_by_key_for_external] syntax error... usage: $0 <raw_json> <key>"
    exit 1
  fi

  get_clean_json_value_for_external "$(echo "$1" | sed -n 's/.*"'$2'":\(\[ *\| \[ *\| *\[\)\([^]]*\).*/\2/p')"
}

get_clean_json_value() {
  if [ "$#" -ne 1 ]; then
    echo "[get_clean_json_value] syntax error... usage: $0 <json_value>"
    exit 1
  fi

  get_clean_string_from_space "$1" | tr -d '"' | tr ',' '\n' | tr -s '\n' ','
}

get_clean_json_value_for_external() {
  if [ "$#" -ne 1 ]; then
    echo "[get_clean_json_value_for_external] syntax error... usage: $0 <json_value>"
    exit 1
  fi

  echo "$1" | tr -d '"' |  tr ',' '\n' | tr -s '\n' ',' | tr -d '[:blank:]' | sed -e 's/, */,/g' -e 's/, *$/ /'
}

get_clean_string_from_space() {
  if [ "$#" -ne 1 ]; then
    echo "[get_clean_string_from_space] syntax error... usage: $0 <string>"
    exit 1
  fi

  echo "$1" | tr -d '[:space:]'
}

read_input() {
  if [ "$#" -ne 4 ]; then
    echo "[read_input] syntax error... usage: $0 <prompt> <is_wrong> <is_confirmation> <check_dir>"
    exit 1
  fi

  if [ "$2" = "true" ]; then
    if [ "$4" = "true" ] && [ "$read_input_value" != "" ] && [ "$read_input_value" != " " ]; then
      printf "%s" "(Directory exist) $1 "
      read -r read_input_value
    else
      printf "%s" "(Wrong) $1 "
      read -r read_input_value
    fi
  else
    printf "%s" "$1 "
    read -r read_input_value
  fi

  if [ "$3" != "true" ]; then
    if [ "$read_input_value" = "" ]; then
      read_input "$1" "true" "$3" "$4"
      return 1
    fi
  fi

  if [ "$4" = "true" ]; then
    if is_dir "$read_input_value"; then
      read_input "$1" true "$3" "$4"
      return 1
    fi
  fi
}

read_project_name() {
  read_input_value=""

  read_input "Enter project name (ex: my-project)... " false false true
}

read_go_module() {
  if is_golang; then
    read_input_value=""

    read_input "Enter go module (ex: my-project, or example.com/user_example/my-project)... " false false false
  fi
}

read_dont_use_venv() {
  if is_python; then
    read_input_value=""

    read_input "We use the virtual environment (venv) by default. Type 'n' and press Enter if you do not wish to use it.... " false true false

    if [ "$read_input_value" = "n" ]; then
      read_input_value=true
    else
      read_input_value=false
    fi
  fi
}

read_confirmation() {
  read_input_value=""

  read_input "Type 'y' and press Enter to confirm. Otherwise, the process will be aborted... " false true false

  if [ "$read_input_value" != "y" ]; then
    exit 1
  fi

  read_input_value=""
}

print_invalid_command_warning() {
  echo "[ERROR] invalid '$command' syntax, please follow:"
  echo "-------------------------------------"
  echo "$(get_language_command_example)"
  echo "or"
  echo "$(get_language_command_example) version"
  echo "or"
  echo "$(get_language_command_example) [command] [value] [source_prefix] [source]"
  echo ""
  echo "command: $(get_command_example)"
  echo "value: $(get_value_example)"
  echo "source_prefix: $(get_source_prefix_example)"
  echo "source: $(get_source_example)"
  echo ""
  echo "Example:"
  echo ""
  echo "$(get_language_command_example) $import_feature_command $(get_feature_example) $import_command_source_prefix git@example.com:user/project.git"
  echo "$(get_language_command_example) $import_domain_command $(get_domain_example) $import_command_source_prefix ../path/to/destination/project"
  echo "$(get_language_command_example) $export_feature_command $(get_feature_example) $export_command_source_prefix ../path/to/destination/project"
  echo "$(get_language_command_example) $export_domain_command $(get_domain_example) $export_command_source_prefix ../path/to/destination/project"
  echo "-------------------------------------"
}

print_version() {
  echo "$name ($language) $version sh$sh_version"
}

validate_import_or_export_command() {
  if is_import_or_export_command_valid; then
    return 0
  fi

  print_invalid_command_warning
  exit 1
}

remove_file() {
  if [ "$#" -ne 1 ]; then
    echo "[remove_file] syntax error... usage: $0 <path>"
    exit 1
  fi

  if ! is_file "$1"; then
    echo "[remove_file] '$1'... does not exist"
    return 1
  fi

  rm -f "$1"

  echo "[remove_file] '$1'... removed"
}

remove_dir() {
  if [ "$#" -ne 1 ]; then
    echo "[remove_dir] syntax error... usage: $0 <path>"
    exit 1
  fi

  if is_file "$1"; then
    echo "[remove_dir] '$1'... is not a directory"
    return 1
  fi

  if ! is_dir "$1"; then
    echo "[remove_dir] '$1'... does not exist"
    return 1
  fi

  rm -rf "$1"

  echo "[remove_dir] '$1'... removed"
}

remove_contents_from_directory() {
  if [ "$#" -ne 2 ]; then
    echo "[remove_contents_from_directory] syntax error... usage: $0 <base_dir> <contents>"
    exit 1
  fi

  if ! is_dir "$1"; then
    echo "[remove_contents_from_directory] '$1'... does not exist"
    return 1
  fi

  if [ "$2" = "" ] || [ "$2" = " " ]; then
    echo "[remove_contents_from_directory] status... nothing to process"
    return 1
  fi

  echo "$2" | tr ',' '\n' | while IFS= read -r _3_item; do
    _3_item=$(echo "$_3_item" | sed -e 's/^[[:space:]]*//' -e 's/[[:space:]]*$//')

    find "$1" -name "$_3_item" -exec rm -rf {} \;
  done

  IFS=$value_delimiter

  echo "[remove_contents_from_directory] items matching '$2' ... removed from '$1'"
}

remove_unwanted_contents_from_directory() {
  if [ "$#" -ne 1 ]; then
    echo "[remove_unwanted_contents_from_directory] syntax error... usage: $0 <base_dir>"
    exit 1
  fi

  remove_contents_from_directory "$1" "__pycache__, node_modules, package-lock.json"
}

remove_path_values() {
  if [ "$#" -ne 1 ]; then
    echo "[remove_path_values] syntax error... usage: $0 <path_value>"
    exit 1
  fi

  if [ "$1" = "" ] || [ "$1" = " " ]; then
    echo "[remove_path_values] status... nothing to process"
    return 1
  fi

  IFS=$value_delimiter

  set -- $(get_clean_string_from_space "$1")

  while [ $# -gt 0 ]; do
    if [ "$1" = "" ] || [ "$1" = " " ]; then
      shift
      continue
    fi
    if is_file "$1"; then
      remove_file "$1"
    elif is_dir "$1"; then
      remove_dir "$1"
    fi
    shift
  done

  IFS=$old_ifs
}

create_dir() {
  if [ "$#" -ne 1 ]; then
    echo "[create_dir] syntax error... usage: $0 <path>"
    exit 1
  fi

  if is_dir "$1"; then
    echo "[create_dir] '$1'... already exist"
    return 1
  fi

  mkdir -p "$1"

  echo "[create_dir] '$1'... created"
}

replace_string_in_file() {
  if [ "$#" -ne 3 ]; then
    echo "[replace_string_in_file] syntax error... usage: $0 <filepath> <old_string> <new_string>"
    exit 1
  fi

  if is_dir "$1"; then
    echo "[replace_string_in_file] '$1'... is not a file"
    return 1
  fi

  if ! is_file "$1"; then
    echo "[replace_string_in_file] '$1'... does not exist"
    return 1
  fi

  sed -i "s/$2/$3/g" "$1"

  echo "[replace_string_in_file] '$2'... replaced with '$3' in '$1'"
}

replace_string_in_directory_files() {
  if [ "$#" -ne 3 ]; then
    echo "[replace_string_in_directory_files] syntax error... usage: $0 <base_dir> <old_string> <new_string>"
    exit 1
  fi

  if is_file "$1"; then
    echo "[replace_string_in_directory_files] '$1'... is not a directory"
    return 1
  fi

  if ! is_dir "$1"; then
    echo "[replace_string_in_directory_files] '$1'... does not exist"
    return 1
  fi

  find "$1" -type f -exec sed -i "s#$2#$3#g" {} +

  echo "[replace_string_in_directory_files] '$2'... replaced with '$3' in all files inside '$1'"
}

replace_go_module_in_file() {
  if [ "$#" -ne 3 ]; then
    echo "[replace_go_module_in_file] syntax error... usage: $0 <filepath> <old_go_module> <new_go_module>"
    exit 1
  fi

  if is_golang; then
    replace_string_in_file "$1" "\"$2" "\"$3"
  fi
}

replace_go_module_in_directory_files() {
  if [ "$#" -ne 3 ]; then
    echo "[replace_go_module_in_directory_files] syntax error... usage: $0 <base_dir> <old_go_module> <new_go_module>"
    exit 1
  fi

  if is_golang; then
    replace_string_in_directory_files "$1" "\"$2" "\"$3"
  fi
}

post_config() {
  if [ "$#" -ne 3 ]; then
    echo "[post_config] syntax error... usage: $0 <base_dir> <go_module> <dont_use_venv>"
    exit 1
  fi

  if is_golang; then
    post_config_golang "$1" "$2"
  elif is_python; then
    post_config_python "$1" "$3"
  elif is_typescript; then
    post_config_typescript "$1"
  fi
}

post_config_golang() {
  if [ "$#" -ne 2 ]; then
    echo "[post_config_golang] syntax error... usage: $0 <base_dir> <go_module>"
    exit 1
  fi

  echo "[post_config_golang] setup... running"

  cd "$1" || return 1

  if [ "$2" != "" ]; then
    go mod init "$2"
  fi

  go mod tidy
  go mod vendor

  cd "$current_dir" || return 1

  echo "[post_config_golang] setup... done"
}

post_config_python() {
  if [ "$#" -ne 2 ]; then
    echo "[post_config_python] syntax error... usage: $0 <base_dir> <dont_use_venv>"
    exit 1
  fi

  echo "[post_config_python] setup... running"

  cd "$1" || return 1

  if ! is_dir "venv"; then
    if [ "$2" = false ]; then
      python3 -m venv venv
      venv/bin/pip install -r requirements.txt
      venv/bin/pip freeze > requirements.txt
    else
      pip install -r requirements.txt
      pip freeze > requirements.txt
    fi
  else
    venv/bin/pip install -r requirements.txt || pip install -r requirements.txt
    venv/bin/pip freeze > requirements.txt || pip freeze > requirements.txt
  fi

  cd "$current_dir" || return 1

  echo "[post_config_python] setup... done"
}

post_config_typescript() {
  if [ "$#" -ne 1 ]; then
    echo "[post_config_typescript] syntax error... usage: $0 <base_dir>"
    exit 1
  fi

  echo "[post_config_typescript] setup... running"

  cd "$1" || return 1

  npm install

  cd "$current_dir" || return 1

  echo "[post_config_typescript] setup... done"
}

install_golang_package() {
  if [ "$#" -lt 1 ]; then
    echo "[install_golang_package] syntax error... usage: $0 <base_dir> <package...>"
    exit 1
  fi

  if is_golang; then
    cd "$1" || return 1

    shift

    go get "$(get_clean_string_from_space "$@")"

    cd "$current_dir" || return 1
  fi
}

install_python_package() {
  if [ "$#" -lt 1 ]; then
    echo "[install_python_package] syntax error... usage: $0 <base_dir> <package>"
    exit 1
  fi

  if is_python; then
    cd "$1" || return 1

    shift

    if ! is_dir "venv"; then
      pip install "$(get_clean_string_from_space "$@")"
    else
      venv/bin/pip install "$(get_clean_string_from_space "$@")" || pip install "$(get_clean_string_from_space "$@")"
    fi

    cd "$current_dir" || return 1
  fi
}

install_typescript_package() {
  if [ "$#" -lt 1 ]; then
    echo "[install_typescript_package] syntax error... usage: $0 <base_dir> <package>"
    exit 1
  fi

  if is_typescript; then
    cd "$1" || return 1

    shift

    npm install "$(get_clean_string_from_space "$@")"

    cd "$current_dir" || return 1
  fi
}

install_dependency_external() {
  if [ "$#" -ne 2 ]; then
    echo "[install_dependency_external] syntax error... usage: $0 <base_dir> <dependency>"
    exit 1
  fi

  if [ "$2" = "" ] || [ "$2" = " " ]; then
    echo "[install_dependency_external] status... nothing to process"
    return 1
  fi

  _8_base_dir="$1"

  IFS=$value_delimiter

  set -- $2

  while [ $# -gt 0 ]; do
    if [ "$1" = "" ] || [ "$1" = " " ]; then
      shift
      continue
    fi
    install_golang_package "$_8_base_dir" $1
    install_python_package "$_8_base_dir" $1
    install_typescript_package "$_8_base_dir" $1
    shift
  done

  IFS=$old_ifs
}

copy_file() {
  if [ "$#" -ne 2 ]; then
    echo "[copy_file] syntax error... usage: $0 <source_path> <destination_path>"
    exit 1
  fi

  if ! is_file "$1"; then
    echo "[copy_file] '$1'... does not exist"
    return 1
  fi

  if is_dir "$2"; then
    echo "[copy_file] '$2'... is not a file"
    return 1
  fi

  _4_destination_dir=$(dirname "$2")

  if ! is_dir "$_4_destination_dir"; then
    create_dir "$_4_destination_dir"
  fi

  cp -n "$1" "$2"

  echo "[copy_file] '$1'... copied to '$2' (existing file will not be overwritten)"
}

copy_contents() {
  if [ "$#" -ne 2 ]; then
    echo "[copy_contents] syntax error... usage: $0 <source_directory> <destination_directory>"
    exit 1
  fi

  if ! is_dir "$1"; then
    echo "[copy_contents] '$1'... does not exist"
    return 1
  fi

  if ! is_dir "$2"; then
    create_dir "$2"
  fi

  cp -rn "$1"/* "$2"

  echo "[copy_contents] '$1'... copied to '$2' (existing contents will not be overwritten)"
}

copy_dependency_feature() {
  if [ "$#" -ne 3 ]; then
    echo "[copy_dependency_feature] syntax error... usage: $0 <dependency> <source_dir> <dest_dir>"
    exit 1
  fi

  if [ "$1" = "" ] || [ "$1" = " " ]; then
    echo "[copy_dependency_feature] status... nothing to process"
    return 1
  fi

  _6_source_dir="$2"
  _6_dest_dir="$3"
  _6_source_go_module=$(get_go_module_from_path "$_6_source_dir")
  _6_dest_go_module=$(get_go_module_from_path "$_6_dest_dir")
  _6_dependency_feature=""
  _6_dependency_domain=""
  _6_dependency_common=""
  _6_dependency_external=""

  IFS=$value_delimiter

  set -- $(get_clean_string_from_space "$1")

  while [ $# -gt 0 ]; do
    if [ "$1" = "" ] || [ "$1" = " " ]; then
      shift
      continue
    fi
    if ! is_dir "$_6_source_dir/$feature_dir/$1"; then
      echo "[copy_dependency_feature] '$1'... feature does not exist in '$_6_source_dir/$feature_dir'"
      shift
      continue
    fi

    if ! is_file "$_6_source_dir/$feature_dir/$1/dependency.json"; then
      echo "[copy_dependency_feature] '$_6_source_dir/$feature_dir/$1'... warning, no 'dependency.json' file found (this may result in missing package errors later)"
    else
      _6_dependency=$(get_raw_json_from_file "$_6_source_dir/$feature_dir/$1/dependency.json")
      _6_dependency=$(get_clean_string_from_space "$_6_dependency")
      _6_dependency_feature="$_6_dependency_feature,$(get_json_value_by_key "$_6_dependency" "$dependency_feature_key")"
      _6_dependency_domain="$_6_dependency_domain,$(get_json_value_by_key "$_6_dependency" "$dependency_domain_key")"
      _6_dependency_common="$_6_dependency_common,$(get_json_value_by_key "$_6_dependency" "$dependency_common_key")"

      _6_dependency=$(get_raw_json_from_file_for_external "$_6_source_dir/$feature_dir/$1/dependency.json")
      _6_dependency_external="$_6_dependency_external,$(get_json_value_by_key_for_external "$_6_dependency" "$dependency_external_key")"
    fi

    copy_contents "$_6_source_dir/$feature_dir/$1" "$_6_dest_dir/$feature_dir/$1"
    replace_go_module_in_directory_files "$_6_dest_dir/$feature_dir/$1" "$_6_source_go_module" "$_6_dest_go_module"
    remove_unwanted_contents_from_directory "$_6_dest_dir/$feature_dir/$1"
    shift
  done

  IFS=$old_ifs

  copy_dependency_domain "$_6_dependency_domain" "$_6_source_dir" "$_6_dest_dir"
  copy_dependency_common "$_6_dependency_common" "$_6_source_dir" "$_6_dest_dir"
  install_dependency_external "$_6_dest_dir" "$_6_dependency_external"
  copy_dependency_feature "$_6_dependency_feature" "$_6_source_dir" "$_6_dest_dir"
}

copy_dependency_domain() {
  if [ "$#" -ne 3 ]; then
    echo "[copy_dependency_domain] syntax error... usage: $0 <dependency> <source_dir> <dest_dir>"
    exit 1
  fi

  if [ "$1" = "" ] || [ "$1" = " " ]; then
    echo "[copy_dependency_domain] status... nothing to process"
    return 1
  fi

  _5_source_dir="$2"
  _5_dest_dir="$3"
  _5_source_go_module=$(get_go_module_from_path "$_5_source_dir")
  _5_dest_go_module=$(get_go_module_from_path "$_5_dest_dir")

  IFS=$value_delimiter

  set -- $(get_clean_string_from_space "$1")

  while [ $# -gt 0 ]; do
    if [ "$1" = "" ] || [ "$1" = " " ]; then
      shift
      continue
    fi
    copy_file "$_5_source_dir/$domain_dir/$1" "$_5_dest_dir/$domain_dir/$1"
    replace_go_module_in_file "$_5_dest_dir/$domain_dir/$1" "$_5_source_go_module" "$_5_dest_go_module"
    shift
  done

  IFS=$old_ifs
}

copy_dependency_common() {
  if [ "$#" -ne 3 ]; then
    echo "[copy_dependency_common] syntax error... usage: $0 <dependency> <source_dir> <dest_dir>"
    exit 1
  fi

  if [ "$1" = "" ] || [ "$1" = " " ]; then
    echo "[copy_dependency_common] status... nothing to process"
    return 1
  fi

  _5_source_dir="$2"
  _5_dest_dir="$3"
  _5_source_go_module=$(get_go_module_from_path "$_5_source_dir")
  _5_dest_go_module=$(get_go_module_from_path "$_5_dest_dir")

  IFS=$value_delimiter

  set -- $(get_clean_string_from_space "$1")

  while [ $# -gt 0 ]; do
    if [ "$1" = "" ] || [ "$1" = " " ]; then
      shift
      continue
    fi
    copy_file "$_5_source_dir/$common_dir/$1" "$_5_dest_dir/$common_dir/$1"
    replace_go_module_in_file "$_5_dest_dir/$common_dir/$1" "$_5_source_go_module" "$_5_dest_go_module"
    shift
  done

  IFS=$old_ifs
}

main
