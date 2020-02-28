import React, { useState } from 'react';
import axios from 'axios';
import API_ENDPOINTS from '../../api';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles(theme => ({
  root: {
    '& > *': {
      margin: theme.spacing(1),
      width: 200,
    },
  },
}));

export default function Topic() {
  const classes = useStyles();
  const [topic, setTopic] = useState('');

  const config = {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
      "Access-Control-Allow-Origin": "*",
    }
  }

  function handleChange(event) {
    setTopic(event.target.value);
  }

  function handleSubmit(e) {
    e.preventDefault();
    axios.post(API_ENDPOINTS.TOPIC_ENDPOINT, topic, config)
      .then(res => console.log('Response: ', res.data));
  }

  return (
    <form className={classes.root} method="POST" action="/" noValidate autoComplete=" off">
      <TextField id="topic" name="topic" type="text" label="Topic" onChange={handleChange} />
      <Button type="submit" onClick={handleSubmit}>Submit</Button>
    </form >
  );
}