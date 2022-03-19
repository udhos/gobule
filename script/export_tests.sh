
export TEST_SAVE=1 ; go test ./parser ; cat parser/tests2.json | jq > parser/tests/tests.json | unset TEST_SAVE
