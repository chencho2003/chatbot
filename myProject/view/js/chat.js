
const chatMessages = document.getElementById('chatMessages');
const userInput = document.getElementById('userInput');
const sendButton = document.getElementById('sendButton');

// Event listener for send button click
sendButton.addEventListener('click', sendMessage);

// Event listener for Enter key press
userInput.addEventListener('keypress', function(event) {
  if (event.key === 'Enter') {
    sendMessage();
  }
});


// Call the function when the page loads

function sendMessage() {
  const question = userInput.value.trim();

  if (question === '') {
    return;
  }

  appendUserMessage( question);

  // Send the message to the backend
  fetch('/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ question })
  })
    .then(function(response) {
      // if (response.ok) {
        return response.json();
      // } else {
        // throw new Error(response.statusText);

      // }
    })
    .then(function(data) {
      // Read the bot's reply from the response object
      const botReply = data.answer 
      const botReply1 = data.error
      console.log(botReply)
      if (typeof botReply === 'undefined') {
        var botReply2 = botReply1
      }else{
        var botReply2 = botReply
      }
  
      // console.log(data.error)
      setTimeout(function() {
        appendBotReply(botReply2);
      }, 600);
    
    })
    .catch(function(error) {
      console.error(error);
    });

  // Clear the input field
  userInput.value = '';
}

// Function to append a user message to the chat messages div
function appendUserMessage(message) {
  const newUserContainer = document.createElement('div');
  newUserContainer.classList.add('message-container');
  newUserContainer.classList.add('user-container');

  const userAvatar = document.createElement('img');
  userAvatar.src = './images/bot.svg';
  userAvatar.classList.add('user-avatar');

  const newUserMessage = document.createElement('div');
  newUserMessage.textContent = message;
  newUserMessage.classList.add('user-message');

  newUserContainer.appendChild(userAvatar);
  newUserContainer.appendChild(newUserMessage);
  chatMessages.appendChild(newUserContainer);
  chatMessages.scrollTop = chatMessages.scrollHeight;
}


// Function to append a bot reply to the chat messages div
function appendBotReply(message) {
  const newBotContainer = document.createElement('div');
  newBotContainer.classList.add('message-container');
  newBotContainer.classList.add('bot-container');

  const botAvatar = document.createElement('img');
  botAvatar.src = './images/bot.svg';
  botAvatar.classList.add('bot-avatar');

  const newBotMessage = document.createElement('div');
  newBotMessage.classList.add('bot-message');
  newBotMessage.textContent = message;

  newBotContainer.appendChild(botAvatar);
  newBotContainer.appendChild(newBotMessage);
  chatMessages.appendChild(newBotContainer);
  chatMessages.scrollTop = chatMessages.scrollHeight;
}

