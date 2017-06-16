#!/usr/bin/env bats

load helpers

function teardown() {
	cleanup_test
}

# PR#59
@test "kpod version test" {
	run kpod version
	echo "$output"
	[ "$status" -eq 0 ]

}

