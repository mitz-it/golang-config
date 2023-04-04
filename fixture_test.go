package config_test

import "os"

var dotenvDummy = []byte(`
MTZ_DUMMY=DUMMY
MTZ_INTERFACE=20
MTZ_SLICE=val1 val2
MTZ_MAP={"key1":"val1", "key2":"val2"}
MTZ_BOOL=true
MTZ_INT=1
MTZ_FLOAT=4.66
MTZ_TIME="2020-04-03T22:50:45Z"
MTZ_DURATION=4h`,
)

var dotenvInvalid = []byte(`
{
  "foo": "bar"
}`,
)

func init() {
	os.Mkdir("./env", 0755)
	os.WriteFile("./env/dummy.env", dotenvDummy, 0644)
	os.WriteFile("./env/invalid.env", dotenvInvalid, 0644)
}
