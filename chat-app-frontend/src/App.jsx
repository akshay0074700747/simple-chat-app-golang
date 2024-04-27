
import React, { useEffect,useState } from 'react'
import SignupForm from './components/signup/signup'
import { Route, Router, Routes, useNavigate } from "react-router-dom";
import HomePage from './components/home/home';
import SigninForm from './components/signin/signin';
import MessageApp from './components/global chat/global_chat';

export default function App() {
  const [islogged, setislogged] = useState(false)
  useEffect(() => {
    fetch("http://localhost:3000",{
      method : "GET",
    }).then((responce)=>{
      if (!responce.ok) {
        console.log("Network responce was not OK")
        return
      }
      return responce.json()
    }).then((data)=>{
      setislogged(data["islogged"])
      console.log(data)
      console.log("jafdhgkahjfl")
    }).catch((error)=>{
      console.error("Error:", error);
    })
  }, [])
  
  return (
   <>
   <Routes>
    <Route exact path = "/" element = {
      islogged ? <HomePage></HomePage> : <SigninForm></SigninForm>
    }></Route>
    <Route path = "/signup" element = {<SignupForm></SignupForm>} ></Route>
    <Route path = "/signin" element = {<SigninForm></SigninForm>} ></Route>
    <Route path="/test" element = {<MessageApp></MessageApp>}></Route>
   </Routes>
   </>
  )
}
