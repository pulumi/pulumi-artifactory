"""A Python Pulumi program"""

import pulumi

import pulumi_artifactory as artifactory

my_local = artifactory.LocalRepository("pylumipus",
            key="pylumipus",
            package_type="npm")