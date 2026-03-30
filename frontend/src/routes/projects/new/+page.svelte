<script lang="ts">
    import { goto } from '$app/navigation';
    import { projectState } from '$lib/stores/project.svelte';
    import { models } from '$lib/wailsjs/go/models';
    import Button from '$lib/components/basic/button.svelte';
    import Input from '$lib/components/form/input.svelte';

    let formData = $state(models.Project.createFrom({
        name: "",
        description: "",
        awsAccount: "",
        awsLocations: []
    }));

    let errors = $state({
        name: "",
        description: "",
        awsAccount: "",
    });

    const validateForm = () => {
        errors = {
            name: "",
            description: "",
            awsAccount: "",
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
        if (!formData.awsAccount.trim()) {
            errors.awsAccount = "AWS account is required.";
        }
        if (formData.awsAccount && !/^\d{12}$/.test(formData.awsAccount.trim())) {
            errors.awsAccount = "AWS account must be a 12-digit number.";
        }
        if (errors.name || errors.description || errors.awsAccount) {
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
            await projectState.createProject(formData);
            goto(`/project/${encodeURIComponent(formData.name)}`);
        } catch (error) {
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
    <h2 class="form-page__subheader">AWS Account Information</h2>
    <Input label="Account" placeholder="Enter AWS account" bind:value={formData.awsAccount} errorMsg={errors.awsAccount} >
        <p>AWS account number must be a 12-digit number.</p>
    </Input>
  </form>
  <Button subtype="add bottom" onClick={handleSubmit}>
    Submit Project
  </Button>
</div>

<style lang="scss">
</style>