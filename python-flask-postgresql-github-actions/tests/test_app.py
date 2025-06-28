import pytest
from app import create_app, db
from app.models import Message

@pytest.fixture
def app():
    app = create_app()
    app.config['TESTING'] = True
    app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://postgres:postgres@localhost:5432/flaskdb_test'
    
    with app.app_context():
        db.create_all()
        yield app
        db.drop_all()

@pytest.fixture
def client(app):
    return app.test_client()

def test_index(client):
    response = client.get('/')
    assert response.status_code == 200
    assert b"Hello, World!" in response.data