<script lang="ts">
    import { configState } from "$lib/stores/config.svelte";
    import Input from "$lib/components/form/input.svelte";
    import Button from "$lib/components/basic/button.svelte";
    import { models } from "$lib/wailsjs/go/models"
    import Select from "$lib/components/form/select.svelte";
    import { onMount } from "svelte";
    import Tooltip from "$lib/components/basic/tooltip.svelte";

    let {awsCredentials= $bindable(), errors, warnings} : 
    {awsCredentials: models.AWSCredentials, errors: {[key: string]: string}, warnings: {[key: string]: string}} = $props();

    let editingAccessKey = $state(false);
    let editingSecretKey = $state(false);

    let initialCredentials = $state({
        accessKey: "",
        secretKey: "",
        regions: [] as string[],
    });

    onMount(() => {
        const data = {...configState.data} as models.AppConfig;
        if (data?.awsCredentials) {
            initialCredentials.accessKey = `${data.awsCredentials.accessKey}`;
            initialCredentials.secretKey = `${data.awsCredentials.secretKey}`;
            initialCredentials.regions = [...data.awsCredentials.regions];
        }

        editingAccessKey = awsCredentials.awsAccessKey == "";
        editingSecretKey = awsCredentials.awsSecretKey == "";
    });

    let displayAccessKey = $derived.by(() => {
        const key = awsCredentials.accessKey;
        if (!key) return "";
        return key.slice(0, 4) + "*****";
    });

    let displaySecretKey = $derived.by(() => {
        const key = awsCredentials.secretKey;
        if (!key) return "";
        return key.slice(0, 4) + "*****";
    });

    const toggleEdit = (field: "access" | "secret") => {
        if (field === "access") {
            editingAccessKey = !editingAccessKey;
        } else {
            editingSecretKey = !editingSecretKey;
        }
    };

    const resetField = (field: "access" | "secret") => {
        if (field === "access") {
            awsCredentials.accessKey = configState.data?.awsCredentials.accessKey || "";
        } else {
            awsCredentials.secretKey = configState.data?.awsCredentials.secretKey || "";
        }
    };

    const resetAll = () => {
        awsCredentials.accessKey = initialCredentials.accessKey;
        awsCredentials.secretKey = initialCredentials.secretKey;
        awsCredentials.regions = [...initialCredentials.regions];
    };

    const resetRegions = () => {
        awsCredentials.regions = [...initialCredentials.regions];
    };

    const awsOptions = {
        "US": [
            { value: "us-east-1", label: "US East (N. Virginia)" },
            { value: "us-east-2", label: "US East (Ohio)" },
            { value: "us-west-1", label: "US West (N. California)" },
            { value: "us-west-2", label: "US West (Oregon)" },
        ],
        "EU": [
            { value: "eu-west-1", label: "EU (Ireland)" },
            { value: "eu-west-2", label: "EU (London)" },
            { value: "eu-west-3", label: "EU (Paris)" },
            { value: "eu-central-1", label: "EU (Frankfurt)" },
        ],
        "AP": [
            { value: "ap-southeast-1", label: "Asia Pacific (Singapore)" },
            { value: "ap-southeast-2", label: "Asia Pacific (Sydney)" },
            { value: "ap-northeast-1", label: "Asia Pacific (Tokyo)" },
            { value: "ap-northeast-2", label: "Asia Pacific (Seoul)" },
        ],
        "SA": [
            { value: "sa-east-1", label: "South America (São Paulo)" },
        ],
        "CA": [
            { value: "ca-central-1", label: "Canada (Central)" },
        ],
        "AF": [
            { value: "af-south-1", label: "Africa (Cape Town)" },
        ],
        "ME": [
            { value: "me-south-1", label: "Middle East (Bahrain)" },
        ],
    }

</script>


{#snippet editBtnGroup(field: "access" | "secret")}
    <div class="btn-group">
        <Button subtype="primary icon small" onClick={() => toggleEdit(field)}>
            &#x270E;
        </Button>
        <Button subtype="error icon small" onClick={() => resetField(field)}>
            &#x21BA;
        </Button>
    </div>
{/snippet}

<div class="aws-credentials">
    <h2 class="form-page__subheader" id="aws-credentials-header">
        AWS Credentials
        <div class="reset" id="aws-credentials-reset">
            <Button subtype="error icon small" onClick={() => resetAll()}>
                &#x21BA;
            </Button>
        </div>
    </h2>
    {#if editingAccessKey}
        <div class="key-row edit">
            <Input type="password" label="Access Key" placeholder="Enter AWS access key" bind:value={awsCredentials.accessKey} errorMsg={errors.awsAccessKey} warningMsg={warnings.awsAccessKey}>
                <p>AWS access key. Required for accessing AWS resources.</p>
            </Input>
            {@render editBtnGroup("access")}
        </div>
    {:else}
        <div class="key-row">
            <Input type="text" label="Access Key" placeholder="Enter AWS access key" disabled={true} value={displayAccessKey} bValidate={true} errorMsg={errors.awsAccessKey} warningMsg={warnings.awsAccessKey}>
                <p>AWS access key. Required for accessing AWS resources.</p>
            </Input>
            {@render editBtnGroup("access")}
        </div>
    {/if}
    {#if editingSecretKey}
        <div class="key-row edit">
            <Input type="password" label="Secret Key" placeholder="Enter AWS secret key" bind:value={awsCredentials.secretKey} errorMsg={errors.awsSecretKey} warningMsg={warnings.awsSecretKey}>
                <p>AWS secret key. Required for authenticating with AWS services.</p>
            </Input>
            {@render editBtnGroup("secret")}
        </div>
    {:else}
        <div class="key-row">
            <Input type="password" label="Secret Key" placeholder="Enter AWS secret key" disabled={true} value={displaySecretKey} bValidate={true} errorMsg={errors.awsSecretKey} warningMsg={warnings.awsSecretKey}>
                <p>AWS secret key. Required for authenticating with AWS services.</p>
            </Input>
            {@render editBtnGroup("secret")}
        </div>
    {/if}
    <h2 class="form-page__subheader" id="aws-regions-header">
        <div class="row title-row">
            <Tooltip position="top">
                <p>AWS regions. Required for specifying the region of AWS resources.</p>
            </Tooltip>
            AWS Regions
        </div>
        <div class="reset" id="aws-regions-reset">
            <Button subtype="add icon small" onClick={() => awsCredentials.regions.push("")}>
                &#x2795;
            </Button>
            <Button subtype="error icon small" onClick={() => resetRegions()}>
                &#x21BA;
            </Button>
        </div>
        <div class="validation-errors">
            {#if errors.awsRegion}
                <p class="error"><span class="icon">&#x20E0;</span>{errors.awsRegion}</p>
            {:else if warnings.awsRegion}
                <p class="warning"><span class="icon">&#x26A0;</span>{warnings.awsRegion}</p>
            {/if}
        </div>
    </h2>
    {#each awsCredentials.regions as region, i}
        <div class="select-row">
            <Select 
                label="" 
                placeholder="Select AWS region" 
                bind:value={awsCredentials.regions[i]} 
                errorMsg={errors.regions?.[i]} 
                warningMsg={warnings.regions?.[i]} 
                options={awsOptions}
                bShowValue={true}
                >
            </Select>
            <div class="btn-group">
                <Button subtype="error icon small" onClick={() => awsCredentials.regions.splice(i, 1)}>
                    &#x2716;
                </Button>
            </div>
        </div>
    {/each}
</div>

<style lang="scss">
    .btn-group {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: calc($padding * -2);
        gap: calc($padding / 2);
    }
    .key-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .select-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: $padding;
    }

    .reset {
        display: flex;
        justify-content: center;
        align-items: center;
        position: absolute;
        top: anchor(top);
        padding: $padding;
        position-area: x-end;
        gap: $padding;
    }

    .form-page__subheader {
        position: relative;
        padding-bottom: calc($padding * 2);
        margin-bottom: calc($padding * 2);
        border-top: 2px solid $clr_light;
        border-bottom: 1px dashed $clr_light;
        margin-top: $padding;
        display: flex;
        flex-direction: column;
        gap: 0;

        .row {
            flex: 1;
            width: 100%;
            text-align: center;
            justify-content: center;
        }

        &-description {
            font-size: 14px;
            color: $clr_light;
            margin-top: $padding;
            font-weight: normal;
            padding: 0;
            margin:0;
        }
    }

    .validation-errors {
        display: flex;
        justify-content: center;
        width: 100%;
        text-align: center;
        height: 14px;
        p {
            text-align: center;
        }
    }
    
    #aws-regions-header {
        anchor-name: --regions-header;
    }
    
    #aws-regions-reset {
        position-anchor: --regions-header;
        width: calc(100% - ($padding));
        justify-content: space-between;
    }
    
    #aws-credentials-header {
        anchor-name: --credentials-header;
    }
    
    #aws-credentials-reset {
        position-anchor: --credentials-header;
        width: calc(100% - ($padding));
        justify-content: flex-end;
    }


</style>