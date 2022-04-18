import * as pulumi from "@pulumi/pulumi";
import * as artifactory from "@pulumi/artifactory";


new artifactory.LocalGenericRepository('pulumipus', {
    key: "pulumipus", // this key can be the name of your repository
})
