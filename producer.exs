Mix.install([
  {:gen_stage, "~> 1.1"}
])

defmodule Producer do
  use GenStage

  def start_link(_args) do
    GenStage.start_link(__MODULE__, [], name: Producer)
  end

  def init(_) do
    IO.puts("Producer started")
    {:producer, 0}
  end

  def handle_demand(demand, counter) do
    IO.puts("Demand is being handled")
    # If the counter is 3 and we ask for 2 items, we will
    # emit the items 3 and 4, and set the state to 5.
    events = Enum.to_list(counter..(counter + demand - 1))
    {:noreply, events, counter + demand}
  end
end

defmodule Main do
  def main do
    children = [
      Producer
    ]

    {:ok, _} = Supervisor.start_link(children, strategy: :one_for_one)
  end
end

Main.main()

# Unless running from IEx, sleep indefinitely so stages keep running
unless IEx.started?() do
  Process.sleep(:infinity)
end
