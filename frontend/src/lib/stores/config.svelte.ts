import { models } from '$lib/wailsjs/go/models';
import { ConfigLoad, ConfigSave, GetAccessKey } from '$lib/wailsjs/go/app/App';
import { snoop } from '$lib/utils';


class ConfigState {
    data = $state<models.AppConfig | null>(null);

    async refresh() {
        let err : string | null = null;
        [this.data, err] = await snoop(ConfigLoad());
        if (err) {
            console.error("Failed to load data:", err);
            throw new Error("Failed to load data: " + err);
        }
        // need try/catch because snoop returns null for string, and TS errors out on that
        try {
            this.data!.awsCredentials.accessKey = await GetAccessKey(this.data!.awsCredentials);
        } catch (err) {
            console.error("Failed to load data:", err);
            throw new Error("Failed to load data: " + err);
        }
    }

    async saveConfig(config: models.AppConfig) {
        const [_, err] = await snoop(ConfigSave(config));
        if (err) {
            console.error("Failed to save config: ", err);
            throw new Error("Failed to save config: " + err);
        }
        await this.refresh();
    }
}

export const configState = new ConfigState();