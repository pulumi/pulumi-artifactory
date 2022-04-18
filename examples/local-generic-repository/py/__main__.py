"""A Python Pulumi program"""

import pulumi

import pulumi_artifactory as artifactory

my_local = artifactory.LocalGenericRepository("pylumipus",
            key="pylumipus")
