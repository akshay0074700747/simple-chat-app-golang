
import React, { useState } from 'react';
import './signup.css';
import { Link } from 'react-router-dom';

const SignupForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    mobile: '',
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
      <h2>Sign Up</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <input
            type="text"
            name="name"
            value={formData.name}
            onChange={handleChange}
            placeholder="Name"
          />
        </div>
        <div className="form-group">
          <input
            type="text"
            name="mobile"
            value={formData.mobile}
            onChange={handleChange}
            placeholder="Mobile"
          />
        </div>
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
        <button type="submit">Sign Up</button>
      </form>
      <p> Already have an Account ? </p>
      <Link to="/signin">Sign In</Link>
    </div>
  );
};

export default SignupForm;
