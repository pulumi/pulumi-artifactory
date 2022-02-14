// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Gradle Repository Resource
//
// Creates a federated Gradle repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewFederatedGradleRepository(ctx, "terraform-federated-test-gradle-repo", &artifactory.FederatedGradleRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-gradle-repo"),
// 			Members: FederatedGradleRepositoryMemberArray{
// 				&FederatedGradleRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-gradle-repo"),
// 				},
// 				&FederatedGradleRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-gradle-repo-2"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FederatedGradleRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedGradleRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                     `pulumi:"notes"`
	PackageType pulumi.StringOutput                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedGradleRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedGradleRepository(ctx *pulumi.Context,
	name string, args *FederatedGradleRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedGradleRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedGradleRepository
	err := ctx.RegisterResource("artifactory:index/federatedGradleRepository:FederatedGradleRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedGradleRepository gets an existing FederatedGradleRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedGradleRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedGradleRepositoryState, opts ...pulumi.ResourceOption) (*FederatedGradleRepository, error) {
	var resource FederatedGradleRepository
	err := ctx.ReadResource("artifactory:index/federatedGradleRepository:FederatedGradleRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedGradleRepository resources.
type federatedGradleRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     []FederatedGradleRepositoryMember `pulumi:"members"`
	Notes       *string                           `pulumi:"notes"`
	PackageType *string                           `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedGradleRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedGradleRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedGradleRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedGradleRepositoryState)(nil)).Elem()
}

type federatedGradleRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members []FederatedGradleRepositoryMember `pulumi:"members"`
	Notes   *string                           `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedGradleRepository resource.
type FederatedGradleRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members FederatedGradleRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedGradleRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedGradleRepositoryArgs)(nil)).Elem()
}

type FederatedGradleRepositoryInput interface {
	pulumi.Input

	ToFederatedGradleRepositoryOutput() FederatedGradleRepositoryOutput
	ToFederatedGradleRepositoryOutputWithContext(ctx context.Context) FederatedGradleRepositoryOutput
}

func (*FederatedGradleRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedGradleRepository)(nil)).Elem()
}

func (i *FederatedGradleRepository) ToFederatedGradleRepositoryOutput() FederatedGradleRepositoryOutput {
	return i.ToFederatedGradleRepositoryOutputWithContext(context.Background())
}

func (i *FederatedGradleRepository) ToFederatedGradleRepositoryOutputWithContext(ctx context.Context) FederatedGradleRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGradleRepositoryOutput)
}

// FederatedGradleRepositoryArrayInput is an input type that accepts FederatedGradleRepositoryArray and FederatedGradleRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedGradleRepositoryArrayInput` via:
//
//          FederatedGradleRepositoryArray{ FederatedGradleRepositoryArgs{...} }
type FederatedGradleRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedGradleRepositoryArrayOutput() FederatedGradleRepositoryArrayOutput
	ToFederatedGradleRepositoryArrayOutputWithContext(context.Context) FederatedGradleRepositoryArrayOutput
}

type FederatedGradleRepositoryArray []FederatedGradleRepositoryInput

func (FederatedGradleRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedGradleRepository)(nil)).Elem()
}

func (i FederatedGradleRepositoryArray) ToFederatedGradleRepositoryArrayOutput() FederatedGradleRepositoryArrayOutput {
	return i.ToFederatedGradleRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedGradleRepositoryArray) ToFederatedGradleRepositoryArrayOutputWithContext(ctx context.Context) FederatedGradleRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGradleRepositoryArrayOutput)
}

// FederatedGradleRepositoryMapInput is an input type that accepts FederatedGradleRepositoryMap and FederatedGradleRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedGradleRepositoryMapInput` via:
//
//          FederatedGradleRepositoryMap{ "key": FederatedGradleRepositoryArgs{...} }
type FederatedGradleRepositoryMapInput interface {
	pulumi.Input

	ToFederatedGradleRepositoryMapOutput() FederatedGradleRepositoryMapOutput
	ToFederatedGradleRepositoryMapOutputWithContext(context.Context) FederatedGradleRepositoryMapOutput
}

type FederatedGradleRepositoryMap map[string]FederatedGradleRepositoryInput

func (FederatedGradleRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedGradleRepository)(nil)).Elem()
}

func (i FederatedGradleRepositoryMap) ToFederatedGradleRepositoryMapOutput() FederatedGradleRepositoryMapOutput {
	return i.ToFederatedGradleRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedGradleRepositoryMap) ToFederatedGradleRepositoryMapOutputWithContext(ctx context.Context) FederatedGradleRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedGradleRepositoryMapOutput)
}

type FederatedGradleRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedGradleRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedGradleRepository)(nil)).Elem()
}

func (o FederatedGradleRepositoryOutput) ToFederatedGradleRepositoryOutput() FederatedGradleRepositoryOutput {
	return o
}

func (o FederatedGradleRepositoryOutput) ToFederatedGradleRepositoryOutputWithContext(ctx context.Context) FederatedGradleRepositoryOutput {
	return o
}

type FederatedGradleRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedGradleRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedGradleRepository)(nil)).Elem()
}

func (o FederatedGradleRepositoryArrayOutput) ToFederatedGradleRepositoryArrayOutput() FederatedGradleRepositoryArrayOutput {
	return o
}

func (o FederatedGradleRepositoryArrayOutput) ToFederatedGradleRepositoryArrayOutputWithContext(ctx context.Context) FederatedGradleRepositoryArrayOutput {
	return o
}

func (o FederatedGradleRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedGradleRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedGradleRepository {
		return vs[0].([]*FederatedGradleRepository)[vs[1].(int)]
	}).(FederatedGradleRepositoryOutput)
}

type FederatedGradleRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedGradleRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedGradleRepository)(nil)).Elem()
}

func (o FederatedGradleRepositoryMapOutput) ToFederatedGradleRepositoryMapOutput() FederatedGradleRepositoryMapOutput {
	return o
}

func (o FederatedGradleRepositoryMapOutput) ToFederatedGradleRepositoryMapOutputWithContext(ctx context.Context) FederatedGradleRepositoryMapOutput {
	return o
}

func (o FederatedGradleRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedGradleRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedGradleRepository {
		return vs[0].(map[string]*FederatedGradleRepository)[vs[1].(string)]
	}).(FederatedGradleRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGradleRepositoryInput)(nil)).Elem(), &FederatedGradleRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGradleRepositoryArrayInput)(nil)).Elem(), FederatedGradleRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedGradleRepositoryMapInput)(nil)).Elem(), FederatedGradleRepositoryMap{})
	pulumi.RegisterOutputType(FederatedGradleRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedGradleRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedGradleRepositoryMapOutput{})
}