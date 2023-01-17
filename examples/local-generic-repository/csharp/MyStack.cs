using Pulumi;
using Pulumi.Artifactory;

class MyStack : Stack
{
    public MyStack()
    {
        var repository = new LocalGenericRepository("csharp-repo", new LocalGenericRepositoryArgs
        {
            Key = "pulumipusdotnet",
            ProjectKey = "default"
        });
    }
}
