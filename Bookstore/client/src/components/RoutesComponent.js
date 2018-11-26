import React, { Component } from 'react';
import { Route, withRouter } from 'react-router-dom';
import Welcome from './Welcome';
import ViewCart from './ViewCart/ViewCart';
import ViewBooks from './Books/ViewBooks';
import BookDetail from './Books/BookDetail';
import Login from './Login/Login'
import Signup from './Login/Signup'

class RoutesComponent extends Component{

    redirectURL = (url) => {
        debugger;
        this.props.history.push(url);

    }

    render(){
        return(
            <div>

                <Route exact path="/" render={() => (
                    <div>
                        <Welcome />
                    </div>
                )}/>

                <Route path="/viewshoppingcart" component={ViewCart}/>
                <Route path="/books" component={ViewBooks}/>
                <Route exact path="/bookdetail/:bookId" component={BookDetail}/>
                <Route path="/login" component={Login}/>
                <Route path="/signup" component={Signup}/>
            </div>
        );
    }




}

export default withRouter(RoutesComponent);