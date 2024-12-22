package test

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"strings"
	"testing"
)

type TestCase struct {
	Name            string
	MethodFunParams []any
	MethodReceiver  any
	PactchFunction  func() *gomonkey.Patches
	ReturnError     bool
}

var err = errors.New("an error occur")

func Test_test(t *testing.T) {
	testCases := []TestCase{
		{
			Name: "flag 为false",
			MethodFunParams: []any{
				11,
				person{},
			},
			MethodReceiver: &method{},
			PactchFunction: func() *gomonkey.Patches {
				return gomonkey.ApplyGlobalVar(&flag, false)
			},
			ReturnError: true,
		},
		{
			Name: "person.name != m.person.name",
			MethodFunParams: []any{
				11,
				person{},
			},
			MethodReceiver: &method{
				person: &person{
					name: "m.person.name",
				},
			},
			PactchFunction: func() *gomonkey.Patches {
				return gomonkey.ApplyFuncReturn(strings.Compare, 1)
			},
			ReturnError: true,
		},
		{
			Name: "add 失败",
			MethodFunParams: []any{
				11,
				person{},
			},
			MethodReceiver: &method{
				person: &person{
					name: "m.person.name",
				},
				data: data{
					age: 1,
				},
			},
			PactchFunction: func() *gomonkey.Patches {
				return gomonkey.ApplyFuncReturn(strings.Compare, 0).
					ApplyFuncReturn(add, 0, err)
			},
			ReturnError: true,
		},
		{
			Name: "Compare 失败",
			MethodFunParams: []any{
				11,
				person{},
			},
			MethodReceiver: &method{
				person: &person{
					name: "m.person.name",
				},
				data: data{
					age: 1,
				},
			},
			PactchFunction: func() *gomonkey.Patches {
				outputs := []gomonkey.OutputCell{
					{Values: gomonkey.Params{0}},
					{Values: gomonkey.Params{1}},
				}
				return gomonkey.ApplyFuncSeq(strings.Compare, outputs).
					ApplyFuncReturn(add, 0, nil)
			},
			ReturnError: true,
		},
		{
			Name: "getOne 失败",
			MethodFunParams: []any{
				11,
				person{},
			},
			MethodReceiver: &method{
				person: &person{
					name: "m.person.name",
				},
				data: data{
					age: 1,
				},
			},
			PactchFunction: func() *gomonkey.Patches {
				outputs := []gomonkey.OutputCell{
					{Values: gomonkey.Params{0}},
					{Values: gomonkey.Params{0}},
				}
				return gomonkey.ApplyFuncSeq(strings.Compare, outputs).
					ApplyFuncReturn(add, 0, nil).
					ApplyPrivateMethod(&group{}, "getOne", func() error {
						return err
					})
			},
			ReturnError: true,
		},
		{
			Name: "test 成功",
			MethodFunParams: []any{
				11,
				person{},
			},
			MethodReceiver: &method{
				person: &person{
					name: "m.person.name",
				},
				data: data{
					age: 1,
				},
			},
			PactchFunction: func() *gomonkey.Patches {
				outputs := []gomonkey.OutputCell{
					{Values: gomonkey.Params{0}},
					{Values: gomonkey.Params{0}},
				}
				return gomonkey.ApplyFuncSeq(strings.Compare, outputs).
					ApplyFuncReturn(add, 0, nil).
					ApplyPrivateMethod(&group{}, "getOne", func() error {
						return nil
					})
			},
			ReturnError: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.PactchFunction != nil {
				patches := testCase.PactchFunction() // 调用Mock函数
				defer patches.Reset()
			}
			age, ok := testCase.MethodFunParams[0].(int)
			if !ok {
				t.Fatalf("test case covert failed, %s", testCase.Name)
			}
			p, ok := testCase.MethodFunParams[1].(person)
			if !ok {
				t.Fatalf("test case covert failed, %s", testCase.Name)
			}
			m, ok := testCase.MethodReceiver.(*method)
			if !ok {
				t.Fatalf("test case covert failed, %s", testCase.Name)
			}
			err := m.test(age, p)
			if (err != nil) != testCase.ReturnError {
				t.Fatalf("test case failed, %s", testCase.Name)
			}
		})
	}
}
