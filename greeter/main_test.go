package greeter_test

import (
	"errors"
	"testing"
)

type testConfig struct {
    args []string
    err error
    config
}

var tests = []testConfig {
    {
        args: []string{"-h"},
        err: nil,
        config: config{printUsage: true, numTimes: 0},
    },
    {
        args: []string{"10"},
        err: nil,
        config: config{printUsage: false, numTimes: 0},
    },
    {
        args: []string{"abcd"},
        err: errors.New("不能转换为数字"),
        config: config{printUsage: false, numTimes: 0},
    },
    {
        args: []string{"l", "bar"},
        err: errors.New("未验证的参数"),
        config: config{printUsage: false, numTimes: 0},
    },
}

func TestParseArgs(t *testing.T) {
    for _, tc := range tests {
        c, err := parseArgs(tc.args)
        if tc.err != nil && err.Error() != tc.err.Error() {
            t.Fatalf("%v\n", tc.err)
        }

        if tc.err == nil && err != nil {
            t.Errorf("%v\n", err)
        }
    }
}
