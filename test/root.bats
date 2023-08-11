#!/usr/bin/env bats

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'

    # get the containing directory of this file
    # use $BATS_TEST_FILENAME instead of ${BASH_SOURCE[0]} or $0,
    # as those will point to the bats executable's location or the preprocessed file respectively
    DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )"
    # make executables in / visible to PATH
    PATH="$DIR/../:$PATH"
}

@test "help command" {
    over --help
}

@test "get command" {
    run over get --config test/testdata/default.yaml
    assert_output "0.1.6"
}

@test "increase command - patch" {
    run over get --config test/testdata/increase.yaml
    assert_output "0.1.6"

    run over up --config test/testdata/increase.yaml --patch --inplace=false
    assert_output "0.1.7"

    run over up --config test/testdata/increase.yaml -p --inplace=false
    assert_output "0.1.7"

    run over up --config test/testdata/increase.yaml -p -i=false
    assert_output "0.1.7"
}

@test "increase command - minor" {
    run over get --config test/testdata/increase.yaml
    assert_output "0.1.6"

    run over up --config test/testdata/increase.yaml --minor --inplace=false
    assert_output "0.2.0"

    run over up --config test/testdata/increase.yaml -m --inplace=false
    assert_output "0.2.0"

    run over up --config test/testdata/increase.yaml -m -i=false
    assert_output "0.2.0"
}

@test "increase command - major" {
    run over get --config test/testdata/increase.yaml
    assert_output "0.1.6"

    run over up --config test/testdata/increase.yaml --major --inplace=false
    assert_output "1.0.0"

    run over up --config test/testdata/increase.yaml -M --inplace=false
    assert_output "1.0.0"

    run over up --config test/testdata/increase.yaml -M -i=false
    assert_output "1.0.0"
}

@test "increase command - files" {
    run over get --config test/testdata/to-change.yaml
    assert_output "0.1.6"

    over up -pi --config test/testdata/to-change.yaml
    run over get --config test/testdata/to-change.yaml
    assert_output "0.1.7"

    over up -mi --config test/testdata/to-change.yaml
    run over get --config test/testdata/to-change.yaml
    assert_output "0.2.0"

    over up -pi --config test/testdata/to-change.yaml
    over up -pi --config test/testdata/to-change.yaml
    over up -pi --config test/testdata/to-change.yaml
    run over get --config test/testdata/to-change.yaml
    assert_output "0.2.3"

    over up -Mi --config test/testdata/to-change.yaml
    run over get --config test/testdata/to-change.yaml
    assert_output "1.0.0"

    echo "package:
  name: to-change
  version: "0.1.6"
  default: true
  files: []
    " > test/testdata/to-change.yaml
}

@test "increase command with V - files" {
    run over get --config test/testdata/to-change-v.yaml
    assert_output "v0.1.6"

    over up -pi --config test/testdata/to-change-v.yaml
    run over get --config test/testdata/to-change-v.yaml
    assert_output "v0.1.7"

    over up -mi --config test/testdata/to-change-v.yaml
    run over get --config test/testdata/to-change-v.yaml
    assert_output "v0.2.0"

    over up -pi --config test/testdata/to-change-v.yaml
    over up -pi --config test/testdata/to-change-v.yaml
    over up -pi --config test/testdata/to-change-v.yaml
    run over get --config test/testdata/to-change-v.yaml
    assert_output "v0.2.3"

    over up -Mi --config test/testdata/to-change-v.yaml
    run over get --config test/testdata/to-change-v.yaml
    assert_output "v1.0.0"

    echo "package:
  name: to-change
  version: "v0.1.6"
  default: true
  files: []
    " > test/testdata/to-change-v.yaml
}