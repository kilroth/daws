import { models } from '$lib/wailsjs/go/models';
import { DataLoad, NewProject, UpdateProject, DeleteProject } from '$lib/wailsjs/go/app/App';
import { snoop } from '$lib/utils';

class ProjectState {
    data = $state<models.AppData | null>(null);

    async refresh() {
        let err : string | null = null;
        [this.data, err] = await snoop(DataLoad());
        if (err) {
            console.error("Failed to load data:", err);
            throw new Error("Failed to load data: " + err);
        }
    }

    async createProject(project: models.Project) {
        const [_, err] = await snoop(NewProject(project));
        if (err) {
            console.error("Failed to create project:", err);
            throw new Error("Failed to create project: " + err);
        }
        await this.refresh();
    }

    async updateProject(oldName: string, project: models.Project) {
        const [_, err] = await snoop(UpdateProject(oldName, project));
        if (err) {
            console.error("Failed to update project:", err);
            throw new Error("Failed to update project: " + err);
        }
        await this.refresh();
    }

    async deleteProject(projectName: string) {
        const [_, err] = await snoop(DeleteProject(projectName));
        if (err) {
            console.error("Failed to delete project:", err);
            throw new Error("Failed to delete project: " + err);
        }
        await this.refresh();
    }
}

export const projectState = new ProjectState();