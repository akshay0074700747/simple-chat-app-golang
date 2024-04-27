
import React, { useState } from 'react';
import './signin.css';
import { Link } from 'react-router-dom';

const SigninForm = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Handle form submission logic here
    console.log(formData);
    fetch("http://localhost:3000/signup",{
        method : "POST",
        body : formData,
    }).then((responce)=>{
        if (!responce.ok) {
            throw new Error("Network response was not ok");
        }
        return responce.json()
    }).then((data)=>{
        console.log("Signup successful:", data["token"]);
    }).catch((error)=>{
        console.error("Error:", error);
    })
  };

  return (
    <div className="signup-form">
      <h2>Sign In</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            placeholder="Email"
          />
        </div>
        <div className="form-group">
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            placeholder="Password"
          />
        </div>
        <button type="submit">Sign In</button>
      </form>
      <p> Don't have an Account ? </p>
      <Link to="/signup">Sign Up</Link>
    </div>
  );
};

export default SigninForm;
