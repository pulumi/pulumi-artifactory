import * as artifactory from "@pulumi/artifactory";

new artifactory.User("node-test-user", {
  admin: false,
  disableUiAccess: false,
  email: "node-test-user@artifactory-terraform.com",
  internalPasswordDisabled: false,
  profileUpdatable: true,
});
