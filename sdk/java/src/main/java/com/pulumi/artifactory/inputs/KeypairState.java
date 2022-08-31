// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KeypairState extends com.pulumi.resources.ResourceArgs {

    public static final KeypairState Empty = new KeypairState();

    /**
     * Will be used as a filename when retrieving the public key via REST API.
     * 
     */
    @Import(name="alias")
    private @Nullable Output<String> alias;

    /**
     * @return Will be used as a filename when retrieving the public key via REST API.
     * 
     */
    public Optional<Output<String>> alias() {
        return Optional.ofNullable(this.alias);
    }

    /**
     * A unique identifier for the Key Pair record.
     * 
     */
    @Import(name="pairName")
    private @Nullable Output<String> pairName;

    /**
     * @return A unique identifier for the Key Pair record.
     * 
     */
    public Optional<Output<String>> pairName() {
        return Optional.ofNullable(this.pairName);
    }

    /**
     * Key Pair type. Supported types - GPG and RSA.
     * 
     */
    @Import(name="pairType")
    private @Nullable Output<String> pairType;

    /**
     * @return Key Pair type. Supported types - GPG and RSA.
     * 
     */
    public Optional<Output<String>> pairType() {
        return Optional.ofNullable(this.pairType);
    }

    /**
     * Passphrase will be used to decrypt the private key. Validated server side.
     * 
     */
    @Import(name="passphrase")
    private @Nullable Output<String> passphrase;

    /**
     * @return Passphrase will be used to decrypt the private key. Validated server side.
     * 
     */
    public Optional<Output<String>> passphrase() {
        return Optional.ofNullable(this.passphrase);
    }

    /**
     * - Private key. PEM format will be validated.
     * 
     */
    @Import(name="privateKey")
    private @Nullable Output<String> privateKey;

    /**
     * @return - Private key. PEM format will be validated.
     * 
     */
    public Optional<Output<String>> privateKey() {
        return Optional.ofNullable(this.privateKey);
    }

    /**
     * Public key. PEM format will be validated.
     * 
     */
    @Import(name="publicKey")
    private @Nullable Output<String> publicKey;

    /**
     * @return Public key. PEM format will be validated.
     * 
     */
    public Optional<Output<String>> publicKey() {
        return Optional.ofNullable(this.publicKey);
    }

    /**
     * Unknown usage. Returned in the json payload and cannot be set.
     * 
     */
    @Import(name="unavailable")
    private @Nullable Output<Boolean> unavailable;

    /**
     * @return Unknown usage. Returned in the json payload and cannot be set.
     * 
     */
    public Optional<Output<Boolean>> unavailable() {
        return Optional.ofNullable(this.unavailable);
    }

    private KeypairState() {}

    private KeypairState(KeypairState $) {
        this.alias = $.alias;
        this.pairName = $.pairName;
        this.pairType = $.pairType;
        this.passphrase = $.passphrase;
        this.privateKey = $.privateKey;
        this.publicKey = $.publicKey;
        this.unavailable = $.unavailable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KeypairState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KeypairState $;

        public Builder() {
            $ = new KeypairState();
        }

        public Builder(KeypairState defaults) {
            $ = new KeypairState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias Will be used as a filename when retrieving the public key via REST API.
         * 
         * @return builder
         * 
         */
        public Builder alias(@Nullable Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias Will be used as a filename when retrieving the public key via REST API.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param pairName A unique identifier for the Key Pair record.
         * 
         * @return builder
         * 
         */
        public Builder pairName(@Nullable Output<String> pairName) {
            $.pairName = pairName;
            return this;
        }

        /**
         * @param pairName A unique identifier for the Key Pair record.
         * 
         * @return builder
         * 
         */
        public Builder pairName(String pairName) {
            return pairName(Output.of(pairName));
        }

        /**
         * @param pairType Key Pair type. Supported types - GPG and RSA.
         * 
         * @return builder
         * 
         */
        public Builder pairType(@Nullable Output<String> pairType) {
            $.pairType = pairType;
            return this;
        }

        /**
         * @param pairType Key Pair type. Supported types - GPG and RSA.
         * 
         * @return builder
         * 
         */
        public Builder pairType(String pairType) {
            return pairType(Output.of(pairType));
        }

        /**
         * @param passphrase Passphrase will be used to decrypt the private key. Validated server side.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(@Nullable Output<String> passphrase) {
            $.passphrase = passphrase;
            return this;
        }

        /**
         * @param passphrase Passphrase will be used to decrypt the private key. Validated server side.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(String passphrase) {
            return passphrase(Output.of(passphrase));
        }

        /**
         * @param privateKey - Private key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(@Nullable Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey - Private key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param publicKey Public key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(@Nullable Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey Public key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        /**
         * @param unavailable Unknown usage. Returned in the json payload and cannot be set.
         * 
         * @return builder
         * 
         */
        public Builder unavailable(@Nullable Output<Boolean> unavailable) {
            $.unavailable = unavailable;
            return this;
        }

        /**
         * @param unavailable Unknown usage. Returned in the json payload and cannot be set.
         * 
         * @return builder
         * 
         */
        public Builder unavailable(Boolean unavailable) {
            return unavailable(Output.of(unavailable));
        }

        public KeypairState build() {
            return $;
        }
    }

}