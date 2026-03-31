<script lang="ts">
    import { goto } from '$app/navigation';
    import { projectState } from '$lib/stores/project.svelte';
    import { models } from '$lib/wailsjs/go/models';
    import Button from '$lib/components/basic/button.svelte';
    import Input from '$lib/components/form/input.svelte';
    import AwsCredentials from '$lib/components/form/awsCredentials.svelte';
    import { onMount } from 'svelte';
    import { configState } from '$lib/stores/config.svelte';

    let formData = $state(models.Project.createFrom({
        name: "",
        description: "",
    }));

    let awsCredentials = $state({
        accessKey: "",
        secretKey: "",
        regions: [] as string[],
    });

    onMount( async () => {
        console.log("Initial form data:", formData);
        await projectState.refresh();
        await configState.refresh();

        const projectData : models.Project = {...projectState.data} as models.Project;
        const configData : models.AppConfig = {...configState.data} as models.AppConfig;
        awsCredentials.regions = projectData?.awsCredentials?.regions.length ? [...projectData.awsCredentials.regions] : [...configData.awsCredentials.regions];
        awsCredentials.accessKey = projectData?.awsCredentials?.accessKey ? `${projectData.awsCredentials.accessKey}` : `${configData.awsCredentials.accessKey}`;
        awsCredentials.secretKey = projectData?.awsCredentials?.secretKey ? `${projectData.awsCredentials.secretKey}` : `${configData.awsCredentials.secretKey}`;
    });

    let errors = $state({
        name: "",
        description: "",
        awsAccessKey: "",
        awsSecretKey: "",
        awsRegion: "",
    });

    let warnings = $state({
        name: "",
        description: "",
        awsAccessKey: "",
        awsSecretKey: "",
        awsRegion: "",
    });

    const validateForm = () => {
        errors = {
            name: "",
            description: "",
            awsAccessKey: "",
            awsSecretKey: "",
            awsRegion: "",
        };
        warnings = {
            name: "",
            description: "",
            awsAccessKey: "",
            awsSecretKey: "",
            awsRegion: "",
        };
        if (!formData.name.trim()) {
            errors.name = "Project name is required.";
        }
        if (formData.name.length > 100) {
            errors.name = "Project name must be less than 100 characters.";
        }
        if (!formData.description.trim()) {
            errors.description = "Project description is required.";
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
        if (errors.name || errors.description || errors.awsAccessKey || errors.awsSecretKey || errors.awsRegion) {
            return false;
        }
        return true;
    };

    async function handleSubmit() {
        /*
        const projectName = prompt("Enter a name for the new project:");
        if (projectName) {
        try {
            await projectState.createProject({ Name: projectName });
            goto(`/project/${encodeURIComponent(projectName)}`);
        } catch (error) {
            alert(`Failed to create project: ${error.message}`);
        }
        }
        */

        if(!validateForm()) {
            return;
        }
        try {
            const outData = {...formData, awsCredentials: {...awsCredentials}} satisfies models.CreateProjectInput;
            const err = await projectState.createProject(outData);
            goto(`/project/${encodeURIComponent(formData.name)}`);
        } catch (error: any) {
            //alert(`Failed to create project: ${error.message}`);
            console.error(`Failed to create project: ${error.message}`);
            // do error states
        }       
    }

</script>

<div class="add-project form-page">
  <h1>Create New Project</h1>
  <form action="" class="form-page__form" >
    <Input label="Name" placeholder="Enter project name" bind:value={formData.name} errorMsg={errors.name}>
        <p>Must be unique and less than 100 characters.</p>
    </Input>
    <Input label="Description" placeholder="Enter project description" bind:value={formData.description} errorMsg={errors.description} >
        <p>Provide a brief description of the project.</p>
    </Input>
    <AwsCredentials bind:awsCredentials {errors} {warnings} />
  </form>
  <Button subtype="add bottom" onClick={handleSubmit}>
    Submit Project
  </Button>
</div>

<style lang="scss">
</style>