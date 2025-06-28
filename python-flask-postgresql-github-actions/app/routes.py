from flask import Blueprint, render_template
from .models import Message

main = Blueprint('main', __name__)

@main.route('/')
def index():
    message = Message.query.first()
    if not message:
        message = Message(content="Hello, World!")
        db.session.add(message)
        db.session.commit()
    return render_template('index.html', message=message.content)