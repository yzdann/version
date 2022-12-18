package version

import "testing"

func TestVersionInfo_FullVersionNumber(t *testing.T) {
	type fields struct {
		Revision  string
		Version   string
		BuildDate string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "With version and revision",
			fields: fields{
				Revision:  "a441b15f",
				Version:   "1.1.1",
				BuildDate: "",
			},
			want: "v1.1.1 (a441b15f)",
		},
		{
			name: "With version, revision, and build date",
			fields: fields{
				Revision:  "a441b15f",
				Version:   "1.1.1",
				BuildDate: "1970-01-01T01:01:01Z",
			},
			want: "v1.1.1 (a441b15f), built 1970-01-01T01:01:01Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VersionInfo{
				Revision:  tt.fields.Revision,
				Version:   tt.fields.Version,
				BuildDate: tt.fields.BuildDate,
			}
			if got := c.FullVersionNumber(); got != tt.want {
				t.Errorf("VersionInfo.FullVersionNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
