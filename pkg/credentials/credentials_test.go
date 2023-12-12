/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2017 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package credentials

import (
	"errors"
	"fmt"
	"testing"
)

type credProvider struct {
	creds   Value
	expired bool
	err     error
}

func (s *credProvider) Retrieve() (Value, error) {
	s.expired = false
	return s.creds, s.err
}
func (s *credProvider) IsExpired() bool {
	return s.expired
}

func TestCredentialsGet(t *testing.T) {
	c := New(&credProvider{
		creds: Value{
			AccessKeyID:     "UXHW",
			SecretAccessKey: "MYSECRET",
			SessionToken:    "",
		},
		expired: true,
	})

	creds, err := c.Get()
	if err != nil {
		t.Fatal(err)
	}
	if "UXHW" != creds.AccessKeyID {
		t.Errorf("Expected \"UXHW\", got %s", creds.AccessKeyID)
	}
	if "MYSECRET" != creds.SecretAccessKey {
		t.Errorf("Expected \"MYSECRET\", got %s", creds.SecretAccessKey)
	}
	if creds.SessionToken != "" {
		t.Errorf("Expected session token to be empty, got %s", creds.SessionToken)
	}
}

func TestCredentialsGetWithError(t *testing.T) {
	c := New(&credProvider{err: errors.New("Custom error")})

	_, err := c.Get()
	if err != nil {
		if err.Error() != "Custom error" {
			t.Errorf("Expected \"Custom error\", got %s", err.Error())
		}
	}
}

func TestWithNginx(t *testing.T) {
	var (
		stsOpts = STSAssumeRoleOptions{
			//AccessKey:       "root",
			//SecretKey:       "minio-openIM",
			AccessKey:       "minioadmin",
			SecretKey:       "minioadmin",
			DurationSeconds: 3600,
		}
		//endpoint = "http://10.10.15.174:10405"
		endpoint = "http://10.10.15.174/openim/minio"
	)

	li, err := NewSTSAssumeRole(endpoint, stsOpts)
	if err != nil {
		fmt.Printf("NewSTSAssumeRole failed, %v\n", err)
		return
	}
	fmt.Printf("NewSTSAssumeRole, %v\n", li)

	v, err := li.Get()
	if err != nil {
		fmt.Printf("li.Get failed, %v\n", err)
		return
	}
	fmt.Printf("li.Get, %v\n", v)
}

func TestUploadFileWithNginx(t *testing.T) {
	var (
		stsOpts = STSAssumeRoleOptions{
			//AccessKey:       "root",
			//SecretKey:       "minio-openIM",
			AccessKey:       "minioadmin",
			SecretKey:       "minioadmin",
			DurationSeconds: 3600,
		}
		//endpoint = "http://10.10.15.174:10405"
		endpoint = "http://10.10.15.174/openim/minio"
	)

	li, err := NewSTSAssumeRole(endpoint, stsOpts)
	if err != nil {
		fmt.Printf("NewSTSAssumeRole failed, %v\n", err)
		return
	}
	fmt.Printf("NewSTSAssumeRole, %v\n", li)

	v, err := li.Get()
	if err != nil {
		fmt.Printf("li.Get failed, %v\n", err)
		return
	}
	fmt.Printf("li.Get, %v\n", v)
}

func TestWithLocal(t *testing.T) {
	var (
		stsOpts = STSAssumeRoleOptions{
			AccessKey:       "minioadmin",
			SecretKey:       "minioadmin",
			DurationSeconds: 3600,
		}
		endpoint = "http://10.40.1.233:9000"
		//endpoint = "http://10.10.15.174/openim/minio"
	)

	li, err := NewSTSAssumeRole(endpoint, stsOpts)
	if err != nil {
		fmt.Printf("NewSTSAssumeRole failed, %v\n", err)
		return
	}
	fmt.Printf("NewSTSAssumeRole, %v\n", li)

	v, err := li.Get()
	if err != nil {
		fmt.Printf("li.Get failed, %v\n", err)
		return
	}
	fmt.Printf("li.Get, %v\n", v)
}
