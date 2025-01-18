import type { MetaFunction } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";

import { Button, InputGroup } from "~/components";

import styles from "./register.module.scss";

export const meta: MetaFunction = () => {
  return [
    { title: "Create Accont - Lopes Drive" },
    {
      name: "description",
      content:
        "Welcome to the lopes drive app, please create your account to start using the Lopes Drive app",
    },
  ];
};

export default function Index() {
  return (
    <div className={styles.container}>
      <main className={styles.card}>
        <h1>Create account</h1>

        <Form className={styles["login-form"]}>
          <InputGroup id="username" label="Username:" />
          <InputGroup id="email" label="Email:" />
          <InputGroup id="password" type="password" label="Password:" />
          <InputGroup
            id="confirm_password"
            type="password"
            label="Confirm Password:"
          />

          <div className={styles["login-form__forgot-password"]}>
            <p>Already have an account?</p> <Link to="../login">Login</Link>
          </div>

          <Button type="submit">Create account</Button>
        </Form>
      </main>

      <small>Copyright &copy; 2025 Matheus Lopes. All rights reserved.</small>
    </div>
  );
}
