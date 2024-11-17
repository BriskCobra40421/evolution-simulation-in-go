from graphviz import Digraph
import random

class EvolutionSimulator:
    def __init__(self, generations=5, output_file='evolution_graph'):
        """
        Initialize the evolution simulator.
        :param generations: Number of generations to simulate.
        :param output_file: Name of the output file (without extension).
        """
        self.graph = Digraph(comment='Evolution Simulation')
        self.output_file = output_file
        self.generations = generations
        self.current_generation = 0
        self.nodes = {}  # Store nodes by generation

    def add_organism(self, organism_id, parent_id=None, mutation=None):
        """
        Add an organism to the graph.
        :param organism_id: Unique ID for the organism.
        :param parent_id: ID of the parent organism (optional).
        :param mutation: Description of mutation/event (optional).
        """
        label = f"{organism_id}\n{mutation}" if mutation else organism_id
        self.graph.node(organism_id, label)
        if parent_id:
            self.graph.edge(parent_id, organism_id, label='Mutation')

    def simulate_generation(self):
        """
        Simulate a single generation of evolution.
        """
        if self.current_generation == 0:
            # Start with a single ancestor
            ancestor_id = 'Ancestor'
            self.add_organism(ancestor_id)
            self.nodes[self.current_generation] = [ancestor_id]
        else:
            new_generation = []
            for parent in self.nodes[self.current_generation - 1]:
                # Each organism can give rise to 1-3 new organisms
                num_offspring = random.randint(1, 3)
                for i in range(num_offspring):
                    mutation = random.choice(['Gene Duplication', 'Point Mutation', 'Chromosomal Rearrangement'])
                    organism_id = f"{parent}.{i}"
                    self.add_organism(organism_id, parent, mutation)
                    new_generation.append(organism_id)
            self.nodes[self.current_generation] = new_generation

    def simulate_evolution(self):
        """
        Simulate multiple generations of evolution.
        """
        for _ in range(self.generations):
            self.simulate_generation()
            self.current_generation += 1

    def render(self, view=True):
        """
        Render the graph to a file and optionally open it.
        :param view: Whether to open the output file after rendering.
        """
        self.graph.render(self.output_file, format='png', view=view)


# Example Usage
if __name__ == "__main__":
    # Create the simulator
    simulator = EvolutionSimulator(generations=5, output_file='evolution_simulation')

    # Simulate evolution
    simulator.simulate_evolution()

    # Render the evolutionary tree
    simulator.render()
