// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LdapGroupSettingArgs extends com.pulumi.resources.ResourceArgs {

    public static final LdapGroupSettingArgs Empty = new LdapGroupSettingArgs();

    /**
     * An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    @Import(name="descriptionAttribute", required=true)
    private Output<String> descriptionAttribute;

    /**
     * @return An attribute on the group entry which denoting the group description. Used when importing groups.
     * 
     */
    public Output<String> descriptionAttribute() {
        return this.descriptionAttribute;
    }

    /**
     * The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    @Import(name="filter", required=true)
    private Output<String> filter;

    /**
     * @return The LDAP filter used to search for group entries. Used for importing groups.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }

    /**
     * A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    @Import(name="groupBaseDn")
    private @Nullable Output<String> groupBaseDn;

    /**
     * @return A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
     * 
     */
    public Optional<Output<String>> groupBaseDn() {
        return Optional.ofNullable(this.groupBaseDn);
    }

    /**
     * A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
     * 
     */
    @Import(name="groupMemberAttribute", required=true)
    private Output<String> groupMemberAttribute;

    /**
     * @return A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
     * 
     */
    public Output<String> groupMemberAttribute() {
        return this.groupMemberAttribute;
    }

    /**
     * Attribute on the group entry denoting the group name. Used when importing groups.
     * 
     */
    @Import(name="groupNameAttribute", required=true)
    private Output<String> groupNameAttribute;

    /**
     * @return Attribute on the group entry denoting the group name. Used when importing groups.
     * 
     */
    public Output<String> groupNameAttribute() {
        return this.groupNameAttribute;
    }

    /**
     * The LDAP setting key you want to use for group retrieval. The value for this field corresponds to &#39;enabledLdap&#39; field of the ldap group setting XML block of system configuration.
     * 
     */
    @Import(name="ldapSettingKey", required=true)
    private Output<String> ldapSettingKey;

    /**
     * @return The LDAP setting key you want to use for group retrieval. The value for this field corresponds to &#39;enabledLdap&#39; field of the ldap group setting XML block of system configuration.
     * 
     */
    public Output<String> ldapSettingKey() {
        return this.ldapSettingKey;
    }

    /**
     * Ldap group setting name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Ldap group setting name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
     * - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
     * - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
     * - HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
     * 
     */
    @Import(name="strategy", required=true)
    private Output<String> strategy;

    /**
     * @return The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
     * - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
     * - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
     * - HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
     * 
     */
    public Output<String> strategy() {
        return this.strategy;
    }

    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
     * 
     */
    @Import(name="subTree")
    private @Nullable Output<Boolean> subTree;

    /**
     * @return When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
     * 
     */
    public Optional<Output<Boolean>> subTree() {
        return Optional.ofNullable(this.subTree);
    }

    private LdapGroupSettingArgs() {}

    private LdapGroupSettingArgs(LdapGroupSettingArgs $) {
        this.descriptionAttribute = $.descriptionAttribute;
        this.filter = $.filter;
        this.groupBaseDn = $.groupBaseDn;
        this.groupMemberAttribute = $.groupMemberAttribute;
        this.groupNameAttribute = $.groupNameAttribute;
        this.ldapSettingKey = $.ldapSettingKey;
        this.name = $.name;
        this.strategy = $.strategy;
        this.subTree = $.subTree;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LdapGroupSettingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LdapGroupSettingArgs $;

        public Builder() {
            $ = new LdapGroupSettingArgs();
        }

        public Builder(LdapGroupSettingArgs defaults) {
            $ = new LdapGroupSettingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param descriptionAttribute An attribute on the group entry which denoting the group description. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder descriptionAttribute(Output<String> descriptionAttribute) {
            $.descriptionAttribute = descriptionAttribute;
            return this;
        }

        /**
         * @param descriptionAttribute An attribute on the group entry which denoting the group description. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder descriptionAttribute(String descriptionAttribute) {
            return descriptionAttribute(Output.of(descriptionAttribute));
        }

        /**
         * @param filter The LDAP filter used to search for group entries. Used for importing groups.
         * 
         * @return builder
         * 
         */
        public Builder filter(Output<String> filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param filter The LDAP filter used to search for group entries. Used for importing groups.
         * 
         * @return builder
         * 
         */
        public Builder filter(String filter) {
            return filter(Output.of(filter));
        }

        /**
         * @param groupBaseDn A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupBaseDn(@Nullable Output<String> groupBaseDn) {
            $.groupBaseDn = groupBaseDn;
            return this;
        }

        /**
         * @param groupBaseDn A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupBaseDn(String groupBaseDn) {
            return groupBaseDn(Output.of(groupBaseDn));
        }

        /**
         * @param groupMemberAttribute A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
         * 
         * @return builder
         * 
         */
        public Builder groupMemberAttribute(Output<String> groupMemberAttribute) {
            $.groupMemberAttribute = groupMemberAttribute;
            return this;
        }

        /**
         * @param groupMemberAttribute A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
         * 
         * @return builder
         * 
         */
        public Builder groupMemberAttribute(String groupMemberAttribute) {
            return groupMemberAttribute(Output.of(groupMemberAttribute));
        }

        /**
         * @param groupNameAttribute Attribute on the group entry denoting the group name. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupNameAttribute(Output<String> groupNameAttribute) {
            $.groupNameAttribute = groupNameAttribute;
            return this;
        }

        /**
         * @param groupNameAttribute Attribute on the group entry denoting the group name. Used when importing groups.
         * 
         * @return builder
         * 
         */
        public Builder groupNameAttribute(String groupNameAttribute) {
            return groupNameAttribute(Output.of(groupNameAttribute));
        }

        /**
         * @param ldapSettingKey The LDAP setting key you want to use for group retrieval. The value for this field corresponds to &#39;enabledLdap&#39; field of the ldap group setting XML block of system configuration.
         * 
         * @return builder
         * 
         */
        public Builder ldapSettingKey(Output<String> ldapSettingKey) {
            $.ldapSettingKey = ldapSettingKey;
            return this;
        }

        /**
         * @param ldapSettingKey The LDAP setting key you want to use for group retrieval. The value for this field corresponds to &#39;enabledLdap&#39; field of the ldap group setting XML block of system configuration.
         * 
         * @return builder
         * 
         */
        public Builder ldapSettingKey(String ldapSettingKey) {
            return ldapSettingKey(Output.of(ldapSettingKey));
        }

        /**
         * @param name Ldap group setting name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Ldap group setting name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param strategy The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
         * - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
         * - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
         * - HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
         * 
         * @return builder
         * 
         */
        public Builder strategy(Output<String> strategy) {
            $.strategy = strategy;
            return this;
        }

        /**
         * @param strategy The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
         * - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
         * - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
         * - HIERARCHICAL: The user&#39;s DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou&#39;s or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
         * 
         * @return builder
         * 
         */
        public Builder strategy(String strategy) {
            return strategy(Output.of(strategy));
        }

        /**
         * @param subTree When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
         * 
         * @return builder
         * 
         */
        public Builder subTree(@Nullable Output<Boolean> subTree) {
            $.subTree = subTree;
            return this;
        }

        /**
         * @param subTree When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
         * 
         * @return builder
         * 
         */
        public Builder subTree(Boolean subTree) {
            return subTree(Output.of(subTree));
        }

        public LdapGroupSettingArgs build() {
            if ($.descriptionAttribute == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingArgs", "descriptionAttribute");
            }
            if ($.filter == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingArgs", "filter");
            }
            if ($.groupMemberAttribute == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingArgs", "groupMemberAttribute");
            }
            if ($.groupNameAttribute == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingArgs", "groupNameAttribute");
            }
            if ($.ldapSettingKey == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingArgs", "ldapSettingKey");
            }
            if ($.strategy == null) {
                throw new MissingRequiredPropertyException("LdapGroupSettingArgs", "strategy");
            }
            return $;
        }
    }

}