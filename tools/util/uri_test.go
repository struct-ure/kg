package util

import (
	"testing"
)

func TestURIFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "file",
			args: args{path: "/foo/bar/root/0.IT/2.APIs and Libraries/0.OpenCV1.json"},
			want: "https://struct-ure.org/kg/it/apis-and-libraries/opencv1",
		},
		{
			name: "directory",
			args: args{path: "/root/0.IT/1.Cloud Computing/0.Amazon Web Services"},
			want: "https://struct-ure.org/kg/it/cloud-computing/amazon-web-services",
		},
		{
			name: "category",
			args: args{path: "/root/0.IT/0.Programming Languages/_categories/imperative.json"},
			want: "https://struct-ure.org/kg/it/programming-languages/categories/imperative",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URIFromPath(tt.args.path); got != tt.want {
				t.Errorf("URIFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRankFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{path: "/root/0.IT/1.Cloud Computing/1.Amazon Web Services"},
			want: 1,
		},
		{
			name: "no rank",
			args: args{path: "/root/0.IT/2.APIs and Libraries/OpenCV1.json"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RankFromPath(tt.args.path); got != tt.want {
				t.Errorf("RankFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParentFromURI(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{uri: "https://foo.com/foo/bar-far"},
			want: "https://foo.com/foo",
		},
		{
			name: "uri char",
			args: args{uri: "https://foo.com/foo/bar/c#"},
			want: "https://foo.com/foo/bar",
		},
		{
			name: "slash uri",
			args: args{uri: "https://foo.com/foo/bar/r(slash)3"},
			want: "https://foo.com/foo/bar",
		},
		{
			name: "dash uri",
			args: args{uri: "https://foo.com/foo/bar/asp-xtrend"},
			want: "https://foo.com/foo/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParentFromURI(tt.args.uri); got != tt.want {
				t.Errorf("ParentFromURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabelFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "file",
			args: args{path: "/foo/bar/root/0.IT/2.APIs and Libraries/0.OpenCV1.json"},
			want: "OpenCV1",
		},
		{
			name: "directory",
			args: args{path: "/root/0.IT/1.Cloud Computing/0.Amazon Web Services/_this.json"},
			want: "Amazon Web Services",
		},
		{
			name: "category",
			args: args{path: "/root/0.IT/0.Programming Languages/_categories/imperative.json"},
			want: "imperative",
		},
		{
			name: "slash uri",
			args: args{path: "/foo/bar/root/0.IT/0.R(slash)3.json"},
			want: "R/3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LabelFromPath(tt.args.path); got != tt.want {
				t.Errorf("LabelFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
