"""A Python Pulumi program"""

import pulumi_artifactory as artifactory

my_local = artifactory.LocalGenericRepository(
    "pylumipus-py", key="pylumipus", project_key="default"
)
