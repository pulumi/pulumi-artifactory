"""A Python Pulumi program"""

import pulumi_artifactory as artifactory

test_user = artifactory.User(
    "python-test-user",
    admin=False,
    disable_ui_access=False,
    email="python-test-user@artifactory-terraform.com",
    internal_password_disabled=False,
    profile_updatable=True
)


