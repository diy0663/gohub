package hash

import "testing"

// 本文件由vscode 自动生成  cmd+shift+p 之后搜索 unit 去生成对应的测试用例
// 在下面的TODO 那里, 加 {}, 之后通过 cmd+shift+p fill struct  填充结构体,之后填入对应的值

func TestBcryptHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "验证密码hash",
			args: args{
				password: "12345678",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BcryptHash(tt.args.password); got != tt.want {
				t.Errorf("BcryptHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBcryptCheck(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "验证是否正确",
			args: args{
				password: "12345678",
				hash:     "$2a$14$CgrbkQyb6hXk2iI6DIhZcuNvOi41vbby.ZUmq3xdxlKSL3MVm3cUC",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BcryptCheck(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("BcryptCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBcryptIsHashed(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BcryptIsHashed(tt.args.str); got != tt.want {
				t.Errorf("BcryptIsHashed() = %v, want %v", got, tt.want)
			}
		})
	}
}
