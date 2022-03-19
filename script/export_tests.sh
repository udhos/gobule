
tmp=tmp.json

export TEST_SAVE=$tmp ; go test ./parser ; cat parser/$tmp | jq > parser/tests/tests.json ; unset TEST_SAVE
