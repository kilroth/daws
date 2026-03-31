<script lang="ts">
    let {
        label, 
        value=$bindable(), 
        placeholder, 
        disabled = false,
        errorMsg = "", 
        display = "column",
        warningMsg = "",
    } : 
    {
        label: string,
        value?: string,
        placeholder: string,
        disabled?: boolean,
        errorMsg?: string,
        display?: "row" | "column",
        warningMsg?: string,
    } = $props();
</script>


<div 
    class="textarea" 
    class:error={!!errorMsg} 
    class:warning={!!warningMsg} 
    class:disabled={disabled}
>
    <div 
        class="label-row" 
        class:row={display === "row"}
        class:column={display === "column"}
    >
        <label for="textarea">{label}</label>
        <textarea id="textarea" placeholder={placeholder} {disabled}>{value}</textarea>
    </div>
    <div class="message-row">
        <div class="spacer"></div>
        {#if errorMsg}
            <div class="error">{errorMsg}</div>
        {:else if warningMsg}
            <div class="warning">{warningMsg}</div>
        {/if}
    </div>
</div>