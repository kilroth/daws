import { models } from '$lib/wailsjs/go/models';
import { ConfigLoad, ConfigSave, GetAWSAccessKey } from '$lib/wailsjs/go/app/App';

class ConfigState {
    data = $state<models.AppConfig | null>(null);

    async refresh() {

        try {
            //this.data = await GetInitialData();
            this.data = await ConfigLoad();
            this.data.awsAccessKey = await GetAWSAccessKey();
        } catch (err) {
            console.error("Failed to load data:", err);
        }
    }

    async saveConfig(config: models.AppConfig) {
        await ConfigSave(config);
        await this.refresh();
    }
}

export const configState = new ConfigState();