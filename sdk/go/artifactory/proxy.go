// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory Proxy resource.
//
// This resource configuration corresponds to 'proxies' config block in system configuration XML
// (REST endpoint: [artifactory/api/system/configuration](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GeneralConfiguration)).
//
// ~>The `Proxy` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewProxy(ctx, "my-proxy", &artifactory.ProxyArgs{
//				Key:             pulumi.String("my-proxy"),
//				Host:            pulumi.String("my-proxy.mycompany.com"),
//				Port:            pulumi.Int(8888),
//				Username:        pulumi.String("user1"),
//				Password:        pulumi.String("password"),
//				NtHost:          pulumi.String("MYCOMPANY.COM"),
//				NtDomain:        pulumi.String("MYCOMPANY"),
//				PlatformDefault: pulumi.Bool(false),
//				RedirectToHosts: pulumi.StringArray{
//					pulumi.String("redirec-host.mycompany.com"),
//				},
//				Services: pulumi.StringArray{
//					pulumi.String("jfrt"),
//					pulumi.String("jfxr"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Current Proxy can be imported using `proxy-key` from Artifactory as the `ID`, e.g.
//
// ```sh
// $ pulumi import artifactory:index/proxy:Proxy my-proxy proxy-key
// ```
type Proxy struct {
	pulumi.CustomResourceState

	// The name of the proxy host.
	Host pulumi.StringOutput `pulumi:"host"`
	// The unique ID of the proxy.
	Key pulumi.StringOutput `pulumi:"key"`
	// The proxy domain/realm name.
	NtDomain pulumi.StringOutput `pulumi:"ntDomain"`
	// The computer name of the machine (the machine connecting to the NTLM proxy).
	NtHost pulumi.StringOutput `pulumi:"ntHost"`
	// The proxy password when authentication credentials are required.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
	PlatformDefault pulumi.BoolOutput `pulumi:"platformDefault"`
	// The proxy port number.
	Port pulumi.IntOutput `pulumi:"port"`
	// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
	RedirectToHosts pulumi.StringArrayOutput `pulumi:"redirectToHosts"`
	// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// The proxy username when authentication credentials are required.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewProxy registers a new resource with the given unique name, arguments, and options.
func NewProxy(ctx *pulumi.Context,
	name string, args *ProxyArgs, opts ...pulumi.ResourceOption) (*Proxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Proxy
	err := ctx.RegisterResource("artifactory:index/proxy:Proxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProxy gets an existing Proxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProxyState, opts ...pulumi.ResourceOption) (*Proxy, error) {
	var resource Proxy
	err := ctx.ReadResource("artifactory:index/proxy:Proxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Proxy resources.
type proxyState struct {
	// The name of the proxy host.
	Host *string `pulumi:"host"`
	// The unique ID of the proxy.
	Key *string `pulumi:"key"`
	// The proxy domain/realm name.
	NtDomain *string `pulumi:"ntDomain"`
	// The computer name of the machine (the machine connecting to the NTLM proxy).
	NtHost *string `pulumi:"ntHost"`
	// The proxy password when authentication credentials are required.
	Password *string `pulumi:"password"`
	// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
	PlatformDefault *bool `pulumi:"platformDefault"`
	// The proxy port number.
	Port *int `pulumi:"port"`
	// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
	RedirectToHosts []string `pulumi:"redirectToHosts"`
	// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
	Services []string `pulumi:"services"`
	// The proxy username when authentication credentials are required.
	Username *string `pulumi:"username"`
}

type ProxyState struct {
	// The name of the proxy host.
	Host pulumi.StringPtrInput
	// The unique ID of the proxy.
	Key pulumi.StringPtrInput
	// The proxy domain/realm name.
	NtDomain pulumi.StringPtrInput
	// The computer name of the machine (the machine connecting to the NTLM proxy).
	NtHost pulumi.StringPtrInput
	// The proxy password when authentication credentials are required.
	Password pulumi.StringPtrInput
	// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
	PlatformDefault pulumi.BoolPtrInput
	// The proxy port number.
	Port pulumi.IntPtrInput
	// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
	RedirectToHosts pulumi.StringArrayInput
	// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
	Services pulumi.StringArrayInput
	// The proxy username when authentication credentials are required.
	Username pulumi.StringPtrInput
}

func (ProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyState)(nil)).Elem()
}

type proxyArgs struct {
	// The name of the proxy host.
	Host string `pulumi:"host"`
	// The unique ID of the proxy.
	Key string `pulumi:"key"`
	// The proxy domain/realm name.
	NtDomain *string `pulumi:"ntDomain"`
	// The computer name of the machine (the machine connecting to the NTLM proxy).
	NtHost *string `pulumi:"ntHost"`
	// The proxy password when authentication credentials are required.
	Password *string `pulumi:"password"`
	// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
	PlatformDefault *bool `pulumi:"platformDefault"`
	// The proxy port number.
	Port int `pulumi:"port"`
	// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
	RedirectToHosts []string `pulumi:"redirectToHosts"`
	// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
	Services []string `pulumi:"services"`
	// The proxy username when authentication credentials are required.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Proxy resource.
type ProxyArgs struct {
	// The name of the proxy host.
	Host pulumi.StringInput
	// The unique ID of the proxy.
	Key pulumi.StringInput
	// The proxy domain/realm name.
	NtDomain pulumi.StringPtrInput
	// The computer name of the machine (the machine connecting to the NTLM proxy).
	NtHost pulumi.StringPtrInput
	// The proxy password when authentication credentials are required.
	Password pulumi.StringPtrInput
	// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
	PlatformDefault pulumi.BoolPtrInput
	// The proxy port number.
	Port pulumi.IntInput
	// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
	RedirectToHosts pulumi.StringArrayInput
	// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
	Services pulumi.StringArrayInput
	// The proxy username when authentication credentials are required.
	Username pulumi.StringPtrInput
}

func (ProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proxyArgs)(nil)).Elem()
}

type ProxyInput interface {
	pulumi.Input

	ToProxyOutput() ProxyOutput
	ToProxyOutputWithContext(ctx context.Context) ProxyOutput
}

func (*Proxy) ElementType() reflect.Type {
	return reflect.TypeOf((**Proxy)(nil)).Elem()
}

func (i *Proxy) ToProxyOutput() ProxyOutput {
	return i.ToProxyOutputWithContext(context.Background())
}

func (i *Proxy) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyOutput)
}

// ProxyArrayInput is an input type that accepts ProxyArray and ProxyArrayOutput values.
// You can construct a concrete instance of `ProxyArrayInput` via:
//
//	ProxyArray{ ProxyArgs{...} }
type ProxyArrayInput interface {
	pulumi.Input

	ToProxyArrayOutput() ProxyArrayOutput
	ToProxyArrayOutputWithContext(context.Context) ProxyArrayOutput
}

type ProxyArray []ProxyInput

func (ProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Proxy)(nil)).Elem()
}

func (i ProxyArray) ToProxyArrayOutput() ProxyArrayOutput {
	return i.ToProxyArrayOutputWithContext(context.Background())
}

func (i ProxyArray) ToProxyArrayOutputWithContext(ctx context.Context) ProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyArrayOutput)
}

// ProxyMapInput is an input type that accepts ProxyMap and ProxyMapOutput values.
// You can construct a concrete instance of `ProxyMapInput` via:
//
//	ProxyMap{ "key": ProxyArgs{...} }
type ProxyMapInput interface {
	pulumi.Input

	ToProxyMapOutput() ProxyMapOutput
	ToProxyMapOutputWithContext(context.Context) ProxyMapOutput
}

type ProxyMap map[string]ProxyInput

func (ProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Proxy)(nil)).Elem()
}

func (i ProxyMap) ToProxyMapOutput() ProxyMapOutput {
	return i.ToProxyMapOutputWithContext(context.Background())
}

func (i ProxyMap) ToProxyMapOutputWithContext(ctx context.Context) ProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProxyMapOutput)
}

type ProxyOutput struct{ *pulumi.OutputState }

func (ProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Proxy)(nil)).Elem()
}

func (o ProxyOutput) ToProxyOutput() ProxyOutput {
	return o
}

func (o ProxyOutput) ToProxyOutputWithContext(ctx context.Context) ProxyOutput {
	return o
}

// The name of the proxy host.
func (o ProxyOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The unique ID of the proxy.
func (o ProxyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The proxy domain/realm name.
func (o ProxyOutput) NtDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.NtDomain }).(pulumi.StringOutput)
}

// The computer name of the machine (the machine connecting to the NTLM proxy).
func (o ProxyOutput) NtHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.NtHost }).(pulumi.StringOutput)
}

// The proxy password when authentication credentials are required.
func (o ProxyOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// When set, this proxy will be the default proxy for new remote repositories and for internal HTTP requests issued by Artifactory. Will also be used as proxy for all other services in the platform (for example: Xray, Distribution, etc).
func (o ProxyOutput) PlatformDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *Proxy) pulumi.BoolOutput { return v.PlatformDefault }).(pulumi.BoolOutput)
}

// The proxy port number.
func (o ProxyOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Proxy) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// An optional list of host names to which this proxy may redirect requests. The credentials defined for the proxy are reused by requests redirected to all of these hosts.
func (o ProxyOutput) RedirectToHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringArrayOutput { return v.RedirectToHosts }).(pulumi.StringArrayOutput)
}

// An optional list of services names to which this proxy be the default of. The options are `jfrt`, `jfmc`, `jfxr`, `jfds`.
func (o ProxyOutput) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringArrayOutput { return v.Services }).(pulumi.StringArrayOutput)
}

// The proxy username when authentication credentials are required.
func (o ProxyOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Proxy) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type ProxyArrayOutput struct{ *pulumi.OutputState }

func (ProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Proxy)(nil)).Elem()
}

func (o ProxyArrayOutput) ToProxyArrayOutput() ProxyArrayOutput {
	return o
}

func (o ProxyArrayOutput) ToProxyArrayOutputWithContext(ctx context.Context) ProxyArrayOutput {
	return o
}

func (o ProxyArrayOutput) Index(i pulumi.IntInput) ProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Proxy {
		return vs[0].([]*Proxy)[vs[1].(int)]
	}).(ProxyOutput)
}

type ProxyMapOutput struct{ *pulumi.OutputState }

func (ProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Proxy)(nil)).Elem()
}

func (o ProxyMapOutput) ToProxyMapOutput() ProxyMapOutput {
	return o
}

func (o ProxyMapOutput) ToProxyMapOutputWithContext(ctx context.Context) ProxyMapOutput {
	return o
}

func (o ProxyMapOutput) MapIndex(k pulumi.StringInput) ProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Proxy {
		return vs[0].(map[string]*Proxy)[vs[1].(string)]
	}).(ProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyInput)(nil)).Elem(), &Proxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyArrayInput)(nil)).Elem(), ProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProxyMapInput)(nil)).Elem(), ProxyMap{})
	pulumi.RegisterOutputType(ProxyOutput{})
	pulumi.RegisterOutputType(ProxyArrayOutput{})
	pulumi.RegisterOutputType(ProxyMapOutput{})
}
