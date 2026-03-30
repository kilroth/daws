<script lang="ts">
    import Button from '$lib/components/basic/button.svelte';
    import Input from '$lib/components/form/input.svelte';
    import { models } from '$lib/wailsjs/go/models';
    import { onMount } from 'svelte';
    import {configState} from '$lib/stores/config.svelte';
    import { OpenFolderDialog } from '$lib/wailsjs/go/app/App';
    import Select from '$lib/components/form/select.svelte';

    let defaultDataPath = "";

    let editingAccessKey = $state(false);
    let editingSecretKey = $state(false);

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

    // use a forms class to handle data and validation instead of individual state variables for each field
    let formData = $state({
        dataPath: "",
        awsAccessKey: "",
        awsSecretKey: "",
        awsRegion: "",
    });

    let errors = $state({
        dataPath: "",
        awsAccessKey: "",
        awsSecretKey: "",
        awsRegion: "",
    });

    let warnings = $state({
        dataPath: "",
        awsAccessKey: "",
        awsSecretKey: "",
        awsRegion: "",
    });

    onMount(async () => {
        configState.refresh();
        console.log("Loaded config:", configState.data);
        const data = models.AppConfig.createFrom(configState.data);

        formData.awsRegion = data.awsRegion;
        formData.awsAccessKey = data.awsAccessKey;
        formData.awsSecretKey = data.awsSecretKey;
        formData.dataPath = data.dataPath;
        defaultDataPath = formData.dataPath;

        editingAccessKey = formData.awsAccessKey == "";
        editingSecretKey = formData.awsSecretKey == "";

    });

    const validate = () => {
        errors = {
            dataPath: "",
            awsAccessKey: "",
            awsSecretKey: "",
            awsRegion: "",
        };
        warnings = {
            dataPath: "",
            awsAccessKey: "",
            awsSecretKey: "",
            awsRegion: "",
        };
        if (!formData.dataPath.trim()) {
            warnings.dataPath = "Using default data path.";
        }
        if (!formData.awsAccessKey.trim()) {
            errors.awsAccessKey = "AWS access key is required.";
        }
        if (formData.awsAccessKey && !/^[A-Za-z0-9]{20}$/.test(formData.awsAccessKey.trim())) {
            errors.awsAccessKey = "AWS access key must be a 20-character string.";
        }
        if (!formData.awsSecretKey.trim()) {
            errors.awsSecretKey = "AWS secret key is required.";
        }
        if (!formData.awsRegion.trim()) {
            errors.awsRegion = "AWS region is required.";
        }
        if (errors.dataPath || errors.awsAccessKey || errors.awsSecretKey || errors.awsRegion) {
            return false;
        }
        return true;
    }

    const handleSubmit = async () => {
        // validate
        if (!validate()) {
            return;
        }

        console.log("Submitting config:", formData);
        await configState.saveConfig({
            dataPath: formData.dataPath,
            awsAccessKey: formData.awsAccessKey,
            awsSecretKey: formData.awsSecretKey,
            awsRegion: formData.awsRegion,
            theme: configState.data?.theme || "light",
        });
        await configState.refresh();
    }

    const handleOpenDirectorSelector = async () => {
        const selectedPath = await OpenFolderDialog();
        if (selectedPath) {
            formData.dataPath = selectedPath;
        }
    }

    
    let displayAccessKey = $derived.by(() => {
        const key = formData.awsAccessKey;
        if (!key) return "";
        return key.slice(0, 4) + "*****";
    });

    let displaySecretKey = $derived.by(() => {
        const key = formData.awsSecretKey;
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
            formData.awsAccessKey = configState.data?.awsAccessKey || "";
        } else {
            formData.awsSecretKey = configState.data?.awsSecretKey || "";
        }
    };

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

<div class="config-page form-page">
    <h1>Configuration Page</h1>
    <form action="" class="config-form form-page__form">
        <div class="row full-width no-gap">
            <Input disabled={true} label="Data Path" placeholder={formData.dataPath} bind:value={formData.dataPath} errorMsg={errors.dataPath} warningMsg={warnings.dataPath}>
                <p>Path where application data is stored. Defaults to <code>%APPDATA%/DockerAWS</code> on Windows and <code>~/.config/DockerAWS</code> on Linux.</p>
            </Input>
            <div class="btn-group">
                <Button subtype="add icon small" onClick={handleOpenDirectorSelector}>
                    &#128449;
                </Button>
                <Button subtype="error icon small" onClick={() => formData.dataPath = defaultDataPath}>
                    &#x21BA;
                </Button>
            </div>
        </div>
        <h2 class="form-page__subheader">AWS Credentials</h2>
        {#if editingAccessKey}
        <div class="key-row edit">
            <Input type="password" label="Access Key" placeholder="Enter AWS access key" bind:value={formData.awsAccessKey} errorMsg={errors.awsAccessKey} warningMsg={warnings.awsAccessKey}>
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
                <Input type="password" label="Secret Key" placeholder="Enter AWS secret key" bind:value={formData.awsSecretKey} errorMsg={errors.awsSecretKey} warningMsg={warnings.awsSecretKey}>
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
        <Select 
            label="Region" 
            placeholder="Select AWS region" 
            bind:value={formData.awsRegion} 
            errorMsg={errors.awsRegion} 
            warningMsg={warnings.awsRegion} 
            options={awsOptions}
            bShowValue={true}
            >
            <p>AWS region. Required for specifying the region of AWS resources.</p>
        </Select>
        
    </form>
    <Button subtype="add bottom" onClick={handleSubmit}>
        Submit Config
    </Button>
</div>

<style lang="scss">
    .btn-group {
        margin-top: calc($padding * -2);
        gap: calc($padding / 2);
    }
    .key-row {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .btn-group {
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>