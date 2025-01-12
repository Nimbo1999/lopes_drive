import type { MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "Lopes Drive" },
    { name: "description", content: "Welcome to the lopes drive app!" },
  ];
};

export default function Index() {
  return (
    <div>
      Landing Index
      <Outlet />
    </div>
  );
}
