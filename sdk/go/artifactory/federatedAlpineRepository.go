// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Alpine Repository Resource
//
// Creates a federated Alpine repository
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
// 		_, err := artifactory.NewFederatedAlpineRepository(ctx, "terraform-federated-test-alpine-repo", &artifactory.FederatedAlpineRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-alpine-repo"),
// 			Members: FederatedAlpineRepositoryMemberArray{
// 				&FederatedAlpineRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-alpine-repo"),
// 				},
// 				&FederatedAlpineRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-alpine-repo-2"),
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
type FederatedAlpineRepository struct {
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
	Members     FederatedAlpineRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                     `pulumi:"notes"`
	PackageType pulumi.StringOutput                        `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewFederatedAlpineRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedAlpineRepository(ctx *pulumi.Context,
	name string, args *FederatedAlpineRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedAlpineRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedAlpineRepository
	err := ctx.RegisterResource("artifactory:index/federatedAlpineRepository:FederatedAlpineRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedAlpineRepository gets an existing FederatedAlpineRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedAlpineRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedAlpineRepositoryState, opts ...pulumi.ResourceOption) (*FederatedAlpineRepository, error) {
	var resource FederatedAlpineRepository
	err := ctx.ReadResource("artifactory:index/federatedAlpineRepository:FederatedAlpineRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedAlpineRepository resources.
type federatedAlpineRepositoryState struct {
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
	Members     []FederatedAlpineRepositoryMember `pulumi:"members"`
	Notes       *string                           `pulumi:"notes"`
	PackageType *string                           `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedAlpineRepositoryState struct {
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
	Members     FederatedAlpineRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedAlpineRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedAlpineRepositoryState)(nil)).Elem()
}

type federatedAlpineRepositoryArgs struct {
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
	Members []FederatedAlpineRepositoryMember `pulumi:"members"`
	Notes   *string                           `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedAlpineRepository resource.
type FederatedAlpineRepositoryArgs struct {
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
	Members FederatedAlpineRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedAlpineRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedAlpineRepositoryArgs)(nil)).Elem()
}

type FederatedAlpineRepositoryInput interface {
	pulumi.Input

	ToFederatedAlpineRepositoryOutput() FederatedAlpineRepositoryOutput
	ToFederatedAlpineRepositoryOutputWithContext(ctx context.Context) FederatedAlpineRepositoryOutput
}

func (*FederatedAlpineRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedAlpineRepository)(nil)).Elem()
}

func (i *FederatedAlpineRepository) ToFederatedAlpineRepositoryOutput() FederatedAlpineRepositoryOutput {
	return i.ToFederatedAlpineRepositoryOutputWithContext(context.Background())
}

func (i *FederatedAlpineRepository) ToFederatedAlpineRepositoryOutputWithContext(ctx context.Context) FederatedAlpineRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedAlpineRepositoryOutput)
}

// FederatedAlpineRepositoryArrayInput is an input type that accepts FederatedAlpineRepositoryArray and FederatedAlpineRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedAlpineRepositoryArrayInput` via:
//
//          FederatedAlpineRepositoryArray{ FederatedAlpineRepositoryArgs{...} }
type FederatedAlpineRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedAlpineRepositoryArrayOutput() FederatedAlpineRepositoryArrayOutput
	ToFederatedAlpineRepositoryArrayOutputWithContext(context.Context) FederatedAlpineRepositoryArrayOutput
}

type FederatedAlpineRepositoryArray []FederatedAlpineRepositoryInput

func (FederatedAlpineRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedAlpineRepository)(nil)).Elem()
}

func (i FederatedAlpineRepositoryArray) ToFederatedAlpineRepositoryArrayOutput() FederatedAlpineRepositoryArrayOutput {
	return i.ToFederatedAlpineRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedAlpineRepositoryArray) ToFederatedAlpineRepositoryArrayOutputWithContext(ctx context.Context) FederatedAlpineRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedAlpineRepositoryArrayOutput)
}

// FederatedAlpineRepositoryMapInput is an input type that accepts FederatedAlpineRepositoryMap and FederatedAlpineRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedAlpineRepositoryMapInput` via:
//
//          FederatedAlpineRepositoryMap{ "key": FederatedAlpineRepositoryArgs{...} }
type FederatedAlpineRepositoryMapInput interface {
	pulumi.Input

	ToFederatedAlpineRepositoryMapOutput() FederatedAlpineRepositoryMapOutput
	ToFederatedAlpineRepositoryMapOutputWithContext(context.Context) FederatedAlpineRepositoryMapOutput
}

type FederatedAlpineRepositoryMap map[string]FederatedAlpineRepositoryInput

func (FederatedAlpineRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedAlpineRepository)(nil)).Elem()
}

func (i FederatedAlpineRepositoryMap) ToFederatedAlpineRepositoryMapOutput() FederatedAlpineRepositoryMapOutput {
	return i.ToFederatedAlpineRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedAlpineRepositoryMap) ToFederatedAlpineRepositoryMapOutputWithContext(ctx context.Context) FederatedAlpineRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedAlpineRepositoryMapOutput)
}

type FederatedAlpineRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedAlpineRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedAlpineRepository)(nil)).Elem()
}

func (o FederatedAlpineRepositoryOutput) ToFederatedAlpineRepositoryOutput() FederatedAlpineRepositoryOutput {
	return o
}

func (o FederatedAlpineRepositoryOutput) ToFederatedAlpineRepositoryOutputWithContext(ctx context.Context) FederatedAlpineRepositoryOutput {
	return o
}

type FederatedAlpineRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedAlpineRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedAlpineRepository)(nil)).Elem()
}

func (o FederatedAlpineRepositoryArrayOutput) ToFederatedAlpineRepositoryArrayOutput() FederatedAlpineRepositoryArrayOutput {
	return o
}

func (o FederatedAlpineRepositoryArrayOutput) ToFederatedAlpineRepositoryArrayOutputWithContext(ctx context.Context) FederatedAlpineRepositoryArrayOutput {
	return o
}

func (o FederatedAlpineRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedAlpineRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedAlpineRepository {
		return vs[0].([]*FederatedAlpineRepository)[vs[1].(int)]
	}).(FederatedAlpineRepositoryOutput)
}

type FederatedAlpineRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedAlpineRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedAlpineRepository)(nil)).Elem()
}

func (o FederatedAlpineRepositoryMapOutput) ToFederatedAlpineRepositoryMapOutput() FederatedAlpineRepositoryMapOutput {
	return o
}

func (o FederatedAlpineRepositoryMapOutput) ToFederatedAlpineRepositoryMapOutputWithContext(ctx context.Context) FederatedAlpineRepositoryMapOutput {
	return o
}

func (o FederatedAlpineRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedAlpineRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedAlpineRepository {
		return vs[0].(map[string]*FederatedAlpineRepository)[vs[1].(string)]
	}).(FederatedAlpineRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedAlpineRepositoryInput)(nil)).Elem(), &FederatedAlpineRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedAlpineRepositoryArrayInput)(nil)).Elem(), FederatedAlpineRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedAlpineRepositoryMapInput)(nil)).Elem(), FederatedAlpineRepositoryMap{})
	pulumi.RegisterOutputType(FederatedAlpineRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedAlpineRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedAlpineRepositoryMapOutput{})
}