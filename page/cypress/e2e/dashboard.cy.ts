describe('Dashboard E2E Tests', () => {
  beforeEach(() => {
    cy.visit('/pages/Dashboard');
  });

  it('renders the header and page structure in initial state', () => {
    cy.get('#title').should('contain.text', 'Dashboard');
    cy.get('.sidebar-nav').should('exist');
  });

  it('handles UI elements and placeholder image', () => {
    cy.get('#user-not-found-img').should('be.visible');
  });

  it('toggles hamburger menu and closes in mobile view', () => {
    cy.viewport(375, 667);

    cy.get('.sidebar').should('not.have.class', 'open');
    cy.get('.hamburger').should('be.visible').click();

    cy.get('.sidebar').should('have.class', 'open');
    cy.get('.overlay').should('exist');

    cy.get('.overlay').click({ force: true });
    cy.get('.sidebar').should('not.have.class', 'open');

    cy.get('.hamburger').click();
    cy.get('.sidebar').should('have.class', 'open');
    cy.get('.close-btn').should('be.visible').click();
    cy.get('.sidebar').should('not.have.class', 'open');
  });
});