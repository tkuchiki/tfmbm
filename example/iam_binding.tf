resource "google_project_iam_binding" "foo" {
  project = "test-project"
  role    = "roles/viewer"
  members = [
    "user:foo1@example.com",
    "user:fooBar@example.com",
  ]
}

resource "google_project_iam_binding" "bar" {
  project = "${var.test_var}"
  role    = "roles/bigquer.dataViewer"
  members = [
    "user:foo.bar.baz@example.com",
    "serviceAccount:baz-bar@example.com",
  ]
}
