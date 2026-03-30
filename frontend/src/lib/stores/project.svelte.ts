import { models } from '$lib/wailsjs/go/models';
import { DataLoad, NewProject, UpdateProject, DeleteProject } from '$lib/wailsjs/go/app/App';

class ProjectState {
    data = $state<models.AppData | null>(null);

    async refresh() {
        try {
            //this.data = await GetInitialData();
            this.data = await DataLoad();
        } catch (err) {
            console.error("Failed to load data:", err);
        }
    }

    async createProject(project: models.Project) {
        await NewProject(project);
        await this.refresh();
    }

    async updateProject(oldName: string, project: models.Project) {
        await UpdateProject(oldName, project);
        await this.refresh();
    }

    async deleteProject(projectName: string) {
        await DeleteProject(projectName);
        await this.refresh();
    }
}

export const projectState = new ProjectState();