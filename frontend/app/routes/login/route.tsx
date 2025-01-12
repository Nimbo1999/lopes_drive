import type { MetaFunction } from "@remix-run/node";
import { Form, Link } from "@remix-run/react";

import { Button, InputGroup } from "~/components";

import styles from "./login.module.scss";

export const meta: MetaFunction = () => {
  return [
    { title: "Login - Lopes Drive" },
    {
      name: "description",
      content:
        "Welcome to the lopes drive app, please login to your account to be able to access your private driver",
    },
  ];
};

export default function Index() {
  return (
    <div className={styles.container}>
      <main className={styles.card}>
        <h1>Login</h1>

        <Form className={styles["login-form"]}>
          <InputGroup id="username" label="Username:" />
          <InputGroup id="password" type="password" label="Password:" />

          <div className={styles["login-form__forgot-password"]}>
            <p>Forgot your password?</p>{" "}
            <Link to="../change-password">Change password</Link>
          </div>

          <Button type="submit">Sign in</Button>
        </Form>
      </main>

      <small>Copyright &copy; 2025 Matheus Lopes. All rights reserved.</small>
    </div>
  );
}
