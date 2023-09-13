using Pulumi;
using Pulumi.Artifactory;

class MyStack : Stack
{
    public MyStack()
    {
        var repository = new User("csharp-test-user", new()
        {
            Admin = false,
            DisableUiAccess = false,
            Email = "csharp-test-user@artifactory-terraform.com",
            InternalPasswordDisabled = true,
            ProfileUpdatable = true,
      });
    }
}
