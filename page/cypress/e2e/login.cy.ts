describe('Registration page tests', () => {
  const REGISTRATION_URL = '/';

  beforeEach(() => {
    cy.visit(REGISTRATION_URL);
  });

  describe('Appearance and element rendering', () => {
    it('displays the header, input field, and registration button', () => {
      cy.contains('h1', 'Registration page').should('be.visible');
      
      cy.get('input[placeholder="Enter your user name here"]')
        .should('be.visible')
        .and('not.be.disabled');

      cy.contains('button', 'Registration')
        .should('be.visible')
        .and('not.be.disabled');
    });
  });

  describe('Validation and error handling', () => {
    it('shows an alert if the name field is empty', () => {
      const alertStub = cy.stub();
      cy.on('window:alert', alertStub);

      cy.contains('button', 'Registration')
        .click()
        .then(() => {
          expect(alertStub).to.be.calledWith('Please add a valid user name.');
        });
    });

    it('handles backend error (e.g., 500 status code)', () => {
      cy.intercept('POST', '**/user/username', {
        statusCode: 500,
      }).as('failedRegistration');

      const alertStub = cy.stub();
      cy.on('window:alert', alertStub);

      cy.get('input[placeholder="Enter your user name here"]').type('TestUser');
      cy.contains('button', 'Registration').click();

      cy.wait('@failedRegistration');
      
      cy.then(() => {
        expect(alertStub).to.be.calledWith(
          Cypress.sinon.match(/Error occured why trying to send your name/)
        );
      });
    });
  });

  describe('Successful registration and state change', () => {
    it('shows loading state and redirects on successful registration', () => {
      cy.intercept('POST', '**/user/username', {
        statusCode: 200,
        body: { ip: '192.168.1.100' },
      }).as('successfulRegistration');

      cy.get('input[placeholder="Enter your user name here"]').type('ValidUser');
      cy.contains('button', 'Registration').click();

      cy.wait('@successfulRegistration');

      cy.url().should('include', '/pages/Dashboard');
    });
  });
});