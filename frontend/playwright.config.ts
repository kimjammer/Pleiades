import { defineConfig, devices } from "@playwright/test"

export default defineConfig({
    webServer: [
        {
            command: "npm run preview",
            port: 4173,
        },
        {
            command: "cd ../backend && go run -tags TEST .",
            port: 8080,
            reuseExistingServer: false, // Must start special testing variant
        },
    ],

    testDir: "e2e",
    use: {
        baseURL: "http://localhost:4173",
    },

    projects: [
        {
            name: "setup",
            use: {
                ...devices["Desktop Chrome"],
                channel: "chromium",
            },
            testMatch: /.*\.setup\.ts/,
        },
        {
            name: "chromium",
            use: {
                ...devices["Desktop Chrome"],
                channel: "chromium",
                // Use prepared auth state.
                storageState: "test-results/.auth/user.json",
            },
            dependencies: ["setup"],
        },
    ],
})
