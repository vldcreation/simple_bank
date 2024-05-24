package util

import (
	"errors"
	"log"
	"os"
	"reflect"
	"strings"
)

func SetEnvValue(tag string, s interface{}) error {
	rt := reflect.ValueOf(s).Elem()
	log.Printf("source type : %+v", rt.Kind())
	if rt.Kind() != reflect.Struct {
		return errors.New("source have bad type")
	}

	typ := rt.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.FieldByName(typ.Field(i).Name)
		log.Printf("tag: %s, f: %s, valid: %+v, canset: %+v, fkind: %+v", tag, f.String(), f.IsValid(), f.CanSet(), f.Kind())
		if f.IsValid() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.
			if f.CanSet() {
				// change value of N
				if f.Kind() == reflect.String {
					v := strings.Split(typ.Field(i).Tag.Get(tag), ",")[0] // use split to ignore tag "options" like omitempty, etc.
					log.Printf("tag: %s, value: %s, env: %s", tag, v, os.Getenv(v))
					f.SetString(os.Getenv(v))
				}
			}
		}
	}

	return nil
}
