describe('app', () => {

  before(() => {
    cy.server()
    cy.route({url: '/api/hello', method: 'GET'}).as('getHelloWorld')
  })

  beforeEach(() => {
    cy.resetDatabase()
  })

  it('lets users view and edit the page banner', () => {
    cy.visit('/')
    cy.wait('@getHelloWorld').should((xhr) => {
      expect(xhr.response!.body).to.deep.contain({
        message: 'You are visitor 0',
      })
    })
    expect(cy.contains('You are visitor 0')).exist

    cy.reload()
    cy.wait('@getHelloWorld').should(xhr => {
      expect(xhr.response!.body).to.deep.contain({
        message: 'You are visitor 1',
      })
    })
    expect(cy.contains('You are visitor 1')).exist

  })
})
