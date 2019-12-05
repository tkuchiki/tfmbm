# tfmbm

tfmbm is the terraform resource migration tool.

## Installation

Download from https://github.com/tkuchiki/tfmbm/releases

## Usage

```console
$ ./tfmbm --help
usage: tfmbm [<flags>] <filepath>

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
    --version  Show application version.

Args:
  <filepath>  Terraform file path
  
```

## Example

```console
$ make build
$ ./tfmbm example/iam_binding.tf

resource "google_project_iam_member" "foo_1_is_viewer" {
  project = "test-project"
  role    = "roles/viewer"
  member  = "user:foo1@example.com"
}

resource "google_project_iam_member" "foo_bar_is_viewer" {
  project = "test-project"
  role    = "roles/viewer"
  member  = "user:fooBar@example.com"
}


resource "google_project_iam_member" "foo_bar_baz_is_bigquer_data_viewer" {
  project = "${var.test_var}"
  role    = "roles/bigquer.dataViewer"
  member  = "user:foo.bar.baz@example.com"
}

resource "google_project_iam_member" "baz_bar_is_bigquer_data_viewer" {
  project = "${var.test_var}"
  role    = "roles/bigquer.dataViewer"
  member  = "serviceAccount:baz-bar@example.com"
}
```