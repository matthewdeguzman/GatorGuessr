describe("Clicks on the map", () => {
    beforeEach(() => {
        cy.visit("http://localhost:4200/login");
        cy.get("[id^=username]").type("Frontend");
        cy.get("[id^=password]").type("FrontendIsNumber1");
        cy.get("[id^=login-button]").click();
    })
    it("Clicks on the map and clicks submit", () => {
        cy.get("[id^=Gmap]").click("center");
        cy.get("[id^=submitButton]").click();
    });
});
describe("Signs out", () => {
    beforeEach(() => {
        cy.visit("http://localhost:4200/login");
        cy.get("[id^=username]").type("Frontend");
        cy.get("[id^=password]").type("FrontendIsNumber1");
        cy.get("[id^=login-button]").click();
    })
    it("Clicks on the signout button", () => {
        cy.contains("Sign Out").click();
        cy.url().should("eq", "http://localhost:4200/home");
    });
});