<script lang="ts">
    import Tooltip from '$lib/components/basic/tooltip.svelte';
    let {
        label, 
        value=$bindable(), 
        placeholder, 
        disabled = false,
        type = "text", 
        display = "row",
        errorMsg = "", 
        warningMsg = "",
        bHidden = false,
        bValidate = true,
        children
    } : 
    {
        label: string,
        value?: string,
        placeholder: string,
        disabled?: boolean,
        type?: string,
        display?: "row" | "column",
        errorMsg?: string,
        warningMsg?: string,
        bHidden?: boolean,
        bValidate?: boolean,
        children?: () => any
    } = $props();

    let displayType = $state(type);

    const showPassword = () => {
        displayType = "text";
    }

    const hidePassword = () => {
        displayType = "password";
    }
</script>

<div class="form-input no-gap" class:row={display === "row"} class:column={display === "column"}>
    <div class="input-row">
        <div class="label-row" class:disabled={disabled && !bValidate}>
            {#if !!children}
                <Tooltip position="top">
                    {@render children()}
                </Tooltip>
            {/if}
            <label for={label} class:error={!!errorMsg} class:warning={!!warningMsg}>{label}</label>
        </div>
        <div class="input-group">
            <input id={label} type={displayType} placeholder={placeholder} bind:value hidden={bHidden} class:error={!!errorMsg} disabled={disabled} />
            {#if type === "password" && !disabled}
                <button class="view" onmousedown={showPassword} onmouseup={hidePassword} onmouseout={hidePassword}>&#x1F441;</button>
            {/if}
        </div>
        
    </div>
    {#if bValidate || (!disabled && !bHidden)}
        <div class="messages-row">
            <div class="spacer"></div>
            {#if errorMsg}
                <p class="error"><span class="icon">&#x20E0;</span>{errorMsg}</p>
            {:else if warningMsg}
                <p class="warning"><span class="icon">&#x26A0;</span>{warningMsg}</p>
            {/if}
        </div>
    {/if}
</div>

<style lang="scss">
    .form-input {        

        &.row {
            align-items: center;
            
        }

        .view {
            height: 36px;
            margin-top: $padding;
            margin-left: calc($padding * -5);
            margin-right: $padding;
            background-color: transparent;
            border: none;
            cursor: pointer;
            font-size: 18px;
            border-radius: 0 $border-radius $border-radius 0;
            color: $clr_dark;
            border-left: 1px solid $clr_dark;
            border-right: 1px solid $clr_dark;

            &:hover {
                color: $clr_light;
                background-color: $clr_dark;
            }
        }

        &.column {
            flex-direction: column;

            label {
                margin-bottom: 0.5em;
            }

            input {
                width: 100%;
            }
        }
    }
</style>