// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory Property Set resource.
 * This resource configuration corresponds to 'propertySets' config block in system configuration XML
 * (REST endpoint: artifactory/api/system/configuration).
 *
 * ~>The `artifactory.PropertySet` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const foo = new artifactory.PropertySet("foo", {
 *     name: "property-set1",
 *     visible: true,
 *     properties: [
 *         {
 *             name: "set1property1",
 *             predefinedValues: [
 *                 {
 *                     name: "passed-QA",
 *                     defaultValue: true,
 *                 },
 *                 {
 *                     name: "failed-QA",
 *                     defaultValue: false,
 *                 },
 *             ],
 *             closedPredefinedValues: true,
 *             multipleChoice: true,
 *         },
 *         {
 *             name: "set1property2",
 *             predefinedValues: [
 *                 {
 *                     name: "passed-QA",
 *                     defaultValue: true,
 *                 },
 *                 {
 *                     name: "failed-QA",
 *                     defaultValue: false,
 *                 },
 *             ],
 *             closedPredefinedValues: false,
 *             multipleChoice: false,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Current Property Set can be imported using `property-set1` as the `ID`, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/propertySet:PropertySet foo property-set1
 * ```
 */
export class PropertySet extends pulumi.CustomResource {
    /**
     * Get an existing PropertySet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PropertySetState, opts?: pulumi.CustomResourceOptions): PropertySet {
        return new PropertySet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/propertySet:PropertySet';

    /**
     * Returns true if the given object is an instance of PropertySet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PropertySet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PropertySet.__pulumiType;
    }

    /**
     * Property set name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of properties that will be part of the property set.
     */
    public readonly properties!: pulumi.Output<outputs.PropertySetProperty[] | undefined>;
    /**
     * Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     */
    public readonly visible!: pulumi.Output<boolean>;

    /**
     * Create a PropertySet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PropertySetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PropertySetArgs | PropertySetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PropertySetState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["properties"] = state ? state.properties : undefined;
            resourceInputs["visible"] = state ? state.visible : undefined;
        } else {
            const args = argsOrState as PropertySetArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["visible"] = args ? args.visible : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PropertySet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PropertySet resources.
 */
export interface PropertySetState {
    /**
     * Property set name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of properties that will be part of the property set.
     */
    properties?: pulumi.Input<pulumi.Input<inputs.PropertySetProperty>[]>;
    /**
     * Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     */
    visible?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a PropertySet resource.
 */
export interface PropertySetArgs {
    /**
     * Property set name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of properties that will be part of the property set.
     */
    properties?: pulumi.Input<pulumi.Input<inputs.PropertySetProperty>[]>;
    /**
     * Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     */
    visible?: pulumi.Input<boolean>;
}