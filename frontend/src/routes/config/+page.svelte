<script lang="ts">
    import Button from '$lib/components/basic/button.svelte';
    import Input from '$lib/components/form/input.svelte';
    import { models } from '$lib/wailsjs/go/models';
    import { onMount } from 'svelte';
    import {configState} from '$lib/stores/config.svelte';
    import { OpenFolderDialog } from '$lib/wailsjs/go/app/App';
    import AwsCredentials from '$lib/components/form/awsCredentials.svelte';

    let defaultDataPath = "";

    // use a forms class to handle data and validation instead of individual state variables for each field
    let formData = $state({
        dataPath: "",
    });

    let awsCredentials = $state({
        accessKey: "",
        secretKey: "",
        regions: [""] as string[],
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
        await configState.refresh();
        console.log("Loaded config:", configState.data);
        const data : models.AppConfig = {...configState.data} as models.AppConfig;
        console.log("Config data:", data);

        formData.dataPath = data.dataPath;
        defaultDataPath = formData.dataPath;
        awsCredentials = {
            accessKey: `${data.awsCredentials.accessKey}`,
            secretKey: `${data.awsCredentials.secretKey}`,
            regions: [...data.awsCredentials.regions],
        };
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
        if (!awsCredentials.accessKey.trim()) {
            errors.awsAccessKey = "AWS access key is required.";
        }
        if (awsCredentials.accessKey && !/^[A-Za-z0-9]{20}$/.test(awsCredentials.accessKey.trim())) {
            errors.awsAccessKey = "AWS access key must be a 20-character string.";
        }
        if (!awsCredentials.secretKey.trim()) {
            errors.awsSecretKey = "AWS secret key is required.";
        }
        if (!awsCredentials.regions.length) {
            errors.awsRegion = "At least one AWS region is required.";
        }
        for (const region of awsCredentials.regions) {
            if (!region.trim()) {
                errors.awsRegion = "AWS region cannot be empty.";
                break;
            }
        }
        console.log("Validation errors:", $state.snapshot(errors));
        console.log("Validation warnings:", $state.snapshot(warnings));
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
            theme: configState.data?.theme || "light",
            awsCredentials,
        });
        await configState.refresh();
    }

    const handleOpenDirectorSelector = async () => {
        const selectedPath = await OpenFolderDialog();
        if (selectedPath) {
            formData.dataPath = selectedPath;
        }
    }

</script>

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
                <Button subtype="error icon small" onClick={() => awsCredentials.awsAccessKey = defaultDataPath}>
                    &#x21BA;
                </Button>
            </div>
        </div>
        <AwsCredentials bind:awsCredentials {errors} {warnings} />
    </form>
    <Button subtype="add bottom" onClick={handleSubmit}>
        Submit Config
    </Button>
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
</style>