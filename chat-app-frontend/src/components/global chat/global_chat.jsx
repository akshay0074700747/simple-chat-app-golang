import React, { useState, useEffect, useRef } from 'react';

const MessageApp = () => {
  const wsRef = useRef(null);
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');

  useEffect(() => {
    const url = "ws://localhost:3000/test";
    const websocket = new WebSocket(url);

    websocket.onopen = () => {
      console.log('WebSocket connection opened');
      wsRef.current = websocket;
    };

    websocket.onmessage = (event) => {
      try {
        const receivedMessage = JSON.parse(event.data);
        setMessages(prevMessages => [...prevMessages, { text: receivedMessage.message, sender: 'bot' }]);
      } catch (error) {
        console.error('Error parsing incoming message:', error);
      }
    };

    websocket.onclose = (event) => {
      console.log('WebSocket connection closed:', event);
    };

    websocket.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    // Clean up WebSocket on component unmount
    return () => {
      if (wsRef.current) {
        wsRef.current.close();
        console.log('WebSocket connection closed during cleanup');
      }
    };
  }, []); // Empty dependency array to ensure useEffect runs only once

  const sendMessage = () => {
    if (wsRef.current && newMessage.trim() !== '') {
      try {
        wsRef.current.send(JSON.stringify({ message: newMessage }));
        setMessages(prevMessages => [...prevMessages, { text: newMessage, sender: 'user' }]);
        setNewMessage('');
      } catch (error) {
        console.error('Error sending message:', error);
      }
    }
  };

  return (
    <div>
      <div style={{ marginBottom: '20px' }}>
        {messages.map((message, index) => (
          <div key={index} style={{ textAlign: message.sender === 'user' ? 'right' : 'left' }}>
            <strong>{message.sender === 'user' ? 'You' : 'Arooo'}:</strong> {message.text}
          </div>
        ))}
      </div>

      <div>
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          placeholder="Type your message..."
        />
        <button onClick={sendMessage}>Send</button>
      </div>
    </div>
  );
};

export default MessageApp;
